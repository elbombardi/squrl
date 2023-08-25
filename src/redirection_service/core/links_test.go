package core

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/elbombardi/squrl/src/db"
	util_mocks "github.com/elbombardi/squrl/src/redirection_service/mocks/util"
	"github.com/elbombardi/squrl/src/redirection_service/util"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setup() (
	*LinksService,
	*db.MockAccountRepository,
	*db.MockLinkRepository,
	*db.MockClickRepository,
	*util.Config,
) {
	accountRepo := new(db.MockAccountRepository)
	linkRepo := new(db.MockLinkRepository)
	clickRepo := new(db.MockClickRepository)
	config := util_mocks.MockConfig()
	return &LinksService{
			AccountRepository: accountRepo,
			LinkRepository:    linkRepo,
			ClickRepository:   clickRepo,
			Config:            config,
			Logger:            util.NewLogger(config),
		},
		accountRepo,
		linkRepo,
		clickRepo,
		config
}

func TestResolveWithNilParams(t *testing.T) {
	s, _, _, _, _ := setup()
	_, err := s.Resolve(nil)
	assert.Error(t, err, "Expected error with nil params")
	assert.Equal(t, ErrBadParams, err.(CoreError).Code, "Expected ErrBadParams error code")
}

func TestResolveWithAccountNotFound(t *testing.T) {
	s, accountRepo, _, _, _ := setup()

	accountRepo.On("GetAccountByPrefix", mock.Anything, "acc").Return(db.Account{}, sql.ErrNoRows)

	params := &ResolveLinkParams{
		AccountPrefix: "acc",
		ShortUrlKey:   "url",
		UserAgent:     "Mozilla/5.0",
		IpAddress:     "127.0.0.1",
	}
	_, err := s.Resolve(params)

	assert.Error(t, err, "Expected error with non-existing account")
	assert.Equal(t, ErrAccountNotFound, err.(CoreError).Code, "Expected ErrAccountNotFound error code")
}

func TestResolveWithAccountDisabled(t *testing.T) {
	s, accountRepo, _, _, _ := setup()

	accountRepo.On("GetAccountByPrefix", mock.Anything, "acc").Return(db.Account{
		Prefix:  "acc",
		Enabled: false,
	}, nil)

	params := &ResolveLinkParams{
		AccountPrefix: "acc",
		ShortUrlKey:   "url",
		UserAgent:     "Mozilla/5.0",
		IpAddress:     "127.0.0.1",
	}
	_, err := s.Resolve(params)

	assert.Error(t, err, "Expected error with non-existing account")
	assert.Equal(t, ErrAccountDisabled, err.(CoreError).Code, "Expected ErrAccountDisabled error code")
}

func TestResolveWithLinkNotFound(t *testing.T) {
	s, accountRepo, linkRepo, _, _ := setup()

	accountRepo.On("GetAccountByPrefix", mock.Anything, "acc").Return(db.Account{
		Prefix:  "acc",
		Enabled: true,
	}, nil)

	linkRepo.On("GetLinkByAccountIDAndShortURLKey", mock.Anything, mock.Anything).Return(db.Link{}, sql.ErrNoRows)

	params := &ResolveLinkParams{
		AccountPrefix: "acc",
		ShortUrlKey:   "url",
		UserAgent:     "Mozilla/5.0",
		IpAddress:     "127.0.0.1",
	}
	_, err := s.Resolve(params)

	assert.Error(t, err, "Expected error with non-existing account")
	assert.Equal(t, ErrLinkNotFound, err.(CoreError).Code, "Expected ErrLinkNotFound error code")
}

func TestResolveWithLinkDisbaled(t *testing.T) {
	s, accountRepo, linkRepo, _, _ := setup()

	accountRepo.On("GetAccountByPrefix", mock.Anything, "acc").Return(db.Account{
		Prefix:  "acc",
		Enabled: true,
	}, nil)

	linkRepo.On("GetLinkByAccountIDAndShortURLKey", mock.Anything, mock.Anything).Return(db.Link{
		Enabled: false,
	}, nil)

	params := &ResolveLinkParams{
		AccountPrefix: "acc",
		ShortUrlKey:   "url",
		UserAgent:     "Mozilla/5.0",
		IpAddress:     "127.0.0.1",
	}
	_, err := s.Resolve(params)

	assert.Error(t, err, "Expected error with non-existing account")
	assert.Equal(t, ErrLinkDisabled, err.(CoreError).Code, "Expected ErrLinkDisabled error code")
}

func TestResolveWithTrackingError(t *testing.T) {
	s, accountRepo, linkRepo, clickRepo, _ := setup()

	accountRepo.On("GetAccountByPrefix", mock.Anything, "acc").Return(db.Account{
		Prefix:  "acc",
		Enabled: true,
	}, nil)

	linkRepo.On("GetLinkByAccountIDAndShortURLKey", mock.Anything, mock.Anything).Return(db.Link{
		Enabled:         true,
		TrackingEnabled: true,
	}, nil)

	clickRepo.On("InsertNewClick", mock.Anything, mock.Anything).Return(errors.New("error"))

	params := &ResolveLinkParams{
		AccountPrefix: "acc",
		ShortUrlKey:   "url",
		UserAgent:     "Mozilla/5.0",
		IpAddress:     "127.0.0.1",
	}
	_, err := s.Resolve(params)

	assert.Error(t, err, "Expected error with non-existing account")
	assert.ErrorContains(t, err, "error", "Expected error to contain 'error'")
}

func TestResolveOk(t *testing.T) {
	s, accountRepo, linkRepo, clickRepo, _ := setup()

	accountRepo.On("GetAccountByPrefix", mock.Anything, "acc").Return(db.Account{
		Prefix:  "acc",
		Enabled: true,
	}, nil)

	linkRepo.On("GetLinkByAccountIDAndShortURLKey", mock.Anything, mock.Anything).Return(db.Link{
		ID:              12,
		Enabled:         true,
		TrackingEnabled: true,
		LongUrl:         "https://example.com",
	}, nil)

	clickRepo.On("InsertNewClick", mock.Anything, mock.Anything).Return(
		func(ctx context.Context, arg db.InsertNewClickParams) error {
			assert.Equal(t, "Mozilla/5.0", arg.UserAgent.String)
			assert.Equal(t, "127.0.0.1", arg.IpAddress.String)
			assert.Equal(t, int32(12), arg.LinkID)
			return nil
		},
	)

	params := &ResolveLinkParams{
		AccountPrefix: "acc",
		ShortUrlKey:   "url",
		UserAgent:     "Mozilla/5.0",
		IpAddress:     "127.0.0.1",
	}
	link, err := s.Resolve(params)

	assert.NoError(t, err, "Expected no error")
	assert.NotNil(t, link, "Expected link to be not nil")
	assert.Equal(t, "https://example.com", link.LongUrl.String(), "Expected link to be https://example.com")
}
