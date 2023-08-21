package core

import (
	"database/sql"
	"errors"
	"testing"

	util_mocks "github.com/elbombardi/squrl/src/api_service/mocks/util"
	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
	db_mocks "github.com/elbombardi/squrl/src/db/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func buildLinkService() (
	*LinksService,
	*db_mocks.MockAccountRepository,
	*db_mocks.MockLinkRepository,
	*util.Config,
) {
	accountRepo := new(db_mocks.MockAccountRepository)
	linkRepo := new(db_mocks.MockLinkRepository)
	config := util_mocks.MockConfig()
	return &LinksService{
			AccountRepository: accountRepo,
			LinkRepository:    linkRepo,
			Config:            config,
			Logger:            util.NewLogger(config),
		},
		accountRepo,
		linkRepo,
		config
}

func TestShortenWithNilUser(t *testing.T) {
	s, _, _, _ := buildLinkService()

	_, err := s.Shorten("http://www.google.com", nil)

	assert.Error(t, err, "Expected error when user is nil")
	assert.Equal(t, ErrUnauthorized, err.(CoreError).Code, "Expected error code to be ERR_UNAUTHORIZED")
}

func TestShortenWithNilParams(t *testing.T) {
	s, _, _, _ := buildLinkService()

	_, err := s.Shorten("", &User{Username: "test"})

	assert.Error(t, err, "Expected error when params are nil")
	assert.Equal(t, ErrBadParams, err.(CoreError).Code, "Expected error code to be ERR_BAD_PARAMS")
}

func TestShortenWithUnknownAccount(t *testing.T) {
	s, accountRepo, _, _ := buildLinkService()

	accountRepo.On("GetAccountByUsername", mock.Anything, "test").Return(db.Account{}, sql.ErrNoRows)

	_, err := s.Shorten("http://www.google.com", &User{Username: "test"})

	assert.Error(t, err, "Expected error when account is unknown")
	assert.Equal(t, ErrAccountNotFound, err.(CoreError).Code, "Expected error code to be ERR_ACCOUNT_NOT_FOUND")
}

func TestShortenErrorWhenLoadingAccount(t *testing.T) {
	s, accountRepo, _, _ := buildLinkService()

	accountRepo.On("GetAccountByUsername", mock.Anything, "test").Return(db.Account{}, errors.New("unexpected_error"))

	_, err := s.Shorten("http://www.google.com", &User{Username: "test"})

	assert.Error(t, err, "Expected error when loading account")
	assert.ErrorContains(t, err, "unexpected_error", "Expected error to contain the error message")
}

func TestShortenWithDisabledAccount(t *testing.T) {
	s, accountRepo, _, _ := buildLinkService()

	accountRepo.On("GetAccountByUsername", mock.Anything, "test").Return(db.Account{Enabled: false}, nil)

	_, err := s.Shorten("http://www.google.com", &User{Username: "test"})

	assert.Error(t, err, "Expected error when account is disabled")
	assert.Equal(t, ErrAccountDisabled, err.(CoreError).Code, "Expected error code to be ERR_ACCOUNT_DISABLED")
}

func TestShortenWithInvalidUrl(t *testing.T) {
	s, _, _, _ := buildLinkService()

	_, err := s.Shorten("invalid_url", &User{Username: "test"})

	assert.Error(t, err, "Expected error when url is invalid")
	assert.Equal(t, ErrBadParams, err.(CoreError).Code, "Expected error code to be ERR_BAD_PARAMS")
}

func TestShortenErrorWhenGeneratingShortUrlKey(t *testing.T) {
	s, accountRepo, linkRepo, _ := buildLinkService()

	accountRepo.On("GetAccountByUsername", mock.Anything, "test").Return(db.Account{Enabled: true}, nil)
	linkRepo.On("CheckShortUrlKeyExists", mock.Anything, mock.Anything).Return(false, errors.New("unexpected_error"))

	_, err := s.Shorten("http://www.google.com", &User{Username: "test"})

	assert.Error(t, err, "Expected error when custom url is invalid")
	assert.ErrorContains(t, err, "unexpected_error", "Expected error to contain the error message")
}

func TestShortenErrorWhenInsertingLink(t *testing.T) {
	s, accountRepo, linkRepo, _ := buildLinkService()

	accountRepo.On("GetAccountByUsername", mock.Anything, "test").Return(db.Account{Enabled: true}, nil)
	linkRepo.On("CheckShortUrlKeyExists", mock.Anything, mock.Anything).Return(false, nil)
	linkRepo.On("InsertNewLink", mock.Anything, mock.Anything).Return(errors.New("unexpected_error"))

	_, err := s.Shorten("http://www.google.com", &User{Username: "test"})

	assert.Error(t, err, "Expected error when custom url is invalid")
	assert.ErrorContains(t, err, "unexpected_error", "Expected error to contain the error message")
}

func TestShortenOk(t *testing.T) {
	s, accountRepo, linkRepo, _ := buildLinkService()

	accountRepo.On("GetAccountByUsername", mock.Anything, "test").Return(db.Account{Prefix: "aaa", Enabled: true}, nil)
	linkRepo.On("CheckShortUrlKeyExists", mock.Anything, mock.Anything).Return(false, nil)
	linkRepo.On("InsertNewLink", mock.Anything, mock.Anything).Return(nil)

	link, err := s.Shorten("http://www.google.com", &User{Username: "test"})

	assert.NoError(t, err, "Expected no error when shortening url")
	assert.NotNil(t, link, "Expected link to be returned")
}
