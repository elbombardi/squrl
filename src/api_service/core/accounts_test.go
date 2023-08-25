package core

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupAccountsService() (
	*AccountsService,
	*db.MockAccountRepository,
	*util.Config,
) {
	accountRepo := new(db.MockAccountRepository)
	config := util.MockConfig()
	return &AccountsService{
			AccountRepository: accountRepo,
			Config:            config,
			Logger:            util.NewLogger(config),
		},
		accountRepo,
		config
}

func TestCreateWithNilUser(t *testing.T) {
	s, _, _ := setupAccountsService()

	params := &CreateAccountParams{}

	_, err := s.Create(params, nil)

	assert.Error(t, err, "Expected error, got nil")
	assert.Equal(t, ErrUnauthorized, err.(CoreError).Code, "Expected error code %d, got %d", ErrUnauthorized, err.(CoreError).Code)
}

func TestCreateWithNonAdminUser(t *testing.T) {
	s, _, _ := setupAccountsService()

	params := &CreateAccountParams{}
	user := &User{Username: "user", IsAdmin: false}

	_, err := s.Create(params, user)

	assert.Error(t, err, "Expected error, got nil")
	assert.Equal(t, ErrUnauthorized, err.(CoreError).Code, "Expected error code %d, got %d", ErrUnauthorized, err.(CoreError).Code)
}

func TestCreateWithoutParams(t *testing.T) {
	s, _, _ := setupAccountsService()

	params := (*CreateAccountParams)(nil)
	user := &User{Username: "user", IsAdmin: true}

	_, err := s.Create(params, user)

	assert.Error(t, err, "Expected error, got nil")
	assert.Equal(t, ErrBadParams, err.(CoreError).Code, "Expected error code %d, got %d", ErrBadParams, err.(CoreError).Code)
}

func TestCeateWithEmptyUsername(t *testing.T) {
	s, _, _ := setupAccountsService()

	user := &User{Username: "admin", IsAdmin: true}
	params := &CreateAccountParams{Username: "", Email: "test@gmail.com"}

	_, err := s.Create(params, user)

	assert.Error(t, err, "Expected error, got nil")
	assert.Equal(t, ErrBadParams, err.(CoreError).Code, "Expected error code %d, got %d", ErrBadParams, err.(CoreError).Code)
}

func TestCreateWithEmptyEmail(t *testing.T) {
	s, _, _ := setupAccountsService()

	user := &User{Username: "admin", IsAdmin: true}
	params := &CreateAccountParams{Username: "account", Email: ""}

	_, err := s.Create(params, user)

	assert.Error(t, err, "Expected error, got nil")
	assert.Equal(t, ErrBadParams, err.(CoreError).Code, "Expected error code %d, got %d", ErrBadParams, err.(CoreError).Code)
}

func TestCeateWithInvalidUsername(t *testing.T) {
	s, _, _ := setupAccountsService()

	user := &User{Username: "admin", IsAdmin: true}
	params := &CreateAccountParams{Username: "####!!!!!!----", Email: "test@gmail.com"}

	_, err := s.Create(params, user)

	assert.Error(t, err, "Expected error, got nil")
	assert.Equal(t, ErrBadParams, err.(CoreError).Code, "Expected error code %d, got %d", ErrBadParams, err.(CoreError).Code)
}

func TestCreateWithInvalidEmail(t *testing.T) {
	s, _, _ := setupAccountsService()

	user := &User{Username: "admin", IsAdmin: true}
	params := &CreateAccountParams{Username: "account", Email: "notanemail"}

	_, err := s.Create(params, user)

	assert.Error(t, err, "Expected error, got nil")
	assert.Equal(t, ErrBadParams, err.(CoreError).Code, "Expected error code %d, got %d", ErrBadParams, err.(CoreError).Code)
}

func TestCreateWithAlreadyExistingUsername(t *testing.T) {
	s, accountRepo, _ := setupAccountsService()

	user := &User{Username: "admin", IsAdmin: true}
	params := &CreateAccountParams{Username: "account", Email: "account@gmail.com"}
	accountRepo.On("CheckUsernameExists", mock.Anything, "account").Return(true, nil)

	_, err := s.Create(params, user)

	assert.Error(t, err, "Expected error, got nil")
	assert.Equal(t, ErrBadParams, err.(CoreError).Code, "Expected error code %d, got %d", ErrBadParams, err.(CoreError).Code)
}

func TestCreateWithErrorWhenTestingUsernameExistance(t *testing.T) {
	s, accountRepo, _ := setupAccountsService()

	user := &User{Username: "admin", IsAdmin: true}
	params := &CreateAccountParams{Username: "account", Email: "account@gmail.com"}
	accountRepo.On("CheckUsernameExists", mock.Anything, "account").Return(false, errors.New("test_error"))

	_, err := s.Create(params, user)

	assert.Error(t, err, "Expected error, got nil")
	assert.ErrorContains(t, err, "test_error", "Expected error to contain 'test_error', got '%s'", err.Error())
}

func TestCreateWithErrorWhileGeneratingPrefix(t *testing.T) {
	s, accountRepo, _ := setupAccountsService()

	user := &User{Username: "admin", IsAdmin: true}
	params := &CreateAccountParams{Username: "account", Email: "account@gmail.com"}
	accountRepo.On("CheckUsernameExists", mock.Anything, "account").Return(false, nil)
	accountRepo.On("CheckPrefixExists", mock.Anything, mock.Anything).Return(false, errors.New("test_error"))

	_, err := s.Create(params, user)

	assert.Error(t, err, "Expected error, got nil")
	assert.ErrorContains(t, err, "test_error", "Expected error to contain 'test_error', got '%s'", err.Error())
}

func TestCreateWithErrorWhilePersistingAccount(t *testing.T) {
	s, accountRepo, _ := setupAccountsService()

	user := &User{Username: "admin", IsAdmin: true}
	params := &CreateAccountParams{Username: "account", Email: "account@gmail.com"}
	accountRepo.On("CheckUsernameExists", mock.Anything, "account").Return(false, nil)
	accountRepo.On("CheckPrefixExists", mock.Anything, mock.Anything).Return(false, nil)
	accountRepo.On("InsertNewAccount", mock.Anything, mock.Anything).Return(errors.New("test_error"))

	_, err := s.Create(params, user)

	assert.Error(t, err, "Expected error, got nil")
	assert.ErrorContains(t, err, "test_error", "Expected error to contain 'test_error', got '%s'", err.Error())
}

func TestCreateOk(t *testing.T) {
	s, accountRepo, _ := setupAccountsService()

	user := &User{Username: "admin", IsAdmin: true}
	params := &CreateAccountParams{Username: "account", Email: "account@gmail.com"}
	accountRepo.On("CheckUsernameExists", mock.Anything, "account").Return(false, nil)
	accountRepo.On("CheckPrefixExists", mock.Anything, mock.Anything).Return(false, nil)
	accountRepo.On("InsertNewAccount", mock.Anything, mock.Anything).Return(nil)

	createdAccount, err := s.Create(params, user)

	assert.NoError(t, err, "Expected no error, got '%s'", err)
	assert.NotNil(t, createdAccount, "Expected account, got nil")
	assert.NotEmpty(t, createdAccount.Password, "Expected password to be set, got empty string")
	assert.NotEmpty(t, createdAccount.Prefix, "Expected prefix to be set, got empty string")
}

func TestUpdateWithNilUser(t *testing.T) {
	s, _, _ := setupAccountsService()

	params := &UpdateAccountParams{}

	_, err := s.Update(params, nil)

	assert.Error(t, err, "Expected error, got nil")
	assert.Equal(t, ErrUnauthorized, err.(CoreError).Code, "Expected error code %d, got %d", ErrUnauthorized, err.(CoreError).Code)
}

func TestUpdateWithNonAdminUser(t *testing.T) {
	s, _, _ := setupAccountsService()

	params := &UpdateAccountParams{}
	user := &User{Username: "user", IsAdmin: false}

	_, err := s.Update(params, user)

	assert.Error(t, err, "Expected error, got nil")
	assert.Equal(t, ErrUnauthorized, err.(CoreError).Code, "Expected error code %d, got %d", ErrUnauthorized, err.(CoreError).Code)
}

func TestUpdateWithoutParams(t *testing.T) {
	s, _, _ := setupAccountsService()

	params := (*UpdateAccountParams)(nil)
	user := &User{Username: "user", IsAdmin: true}

	_, err := s.Update(params, user)

	assert.Error(t, err, "Expected error, got nil")
	assert.Equal(t, ErrBadParams, err.(CoreError).Code, "Expected error code %d, got %d", ErrBadParams, err.(CoreError).Code)
}

func TestUpdateWithNonExistentAccount(t *testing.T) {
	s, accountRepo, _ := setupAccountsService()

	user := &User{Username: "admin", IsAdmin: true}
	params := &UpdateAccountParams{Username: "account1", Enabled: true}
	accountRepo.On("GetAccountByUsername", mock.Anything, "account1").Return(db.Account{}, sql.ErrNoRows)

	_, err := s.Update(params, user)

	assert.Error(t, err, "Expected error, got nil")
	assert.Equal(t, ErrAccountNotFound, err.(CoreError).Code, "Expected error code %d, got %d", ErrAccountNotFound, err.(CoreError).Code)
}

func TestUpdateWithErrorWhileCheckingAccount(t *testing.T) {
	s, accountRepo, _ := setupAccountsService()

	user := &User{Username: "admin", IsAdmin: true}
	params := &UpdateAccountParams{Username: "account1", Enabled: true}
	accountRepo.On("GetAccountByUsername", mock.Anything, "account1").Return(db.Account{}, errors.New("test_error"))

	_, err := s.Update(params, user)

	assert.Error(t, err, "Expected error, got nil")
	assert.ErrorContains(t, err, "test_error", "Expected error to contain 'test_error', got '%s'", err.Error())
}

func TestUpdateWithErrorWhileUpdating(t *testing.T) {
	s, accountRepo, _ := setupAccountsService()

	user := &User{Username: "admin", IsAdmin: true}
	params := &UpdateAccountParams{Username: "account1", Enabled: true}
	accountRepo.On("GetAccountByUsername", mock.Anything, "account1").Return(db.Account{}, nil)
	accountRepo.On("UpdateAccountStatusByUsername", mock.Anything, mock.Anything).Return(errors.New("test_error"))

	_, err := s.Update(params, user)

	assert.Error(t, err, "Expected error, got nil")
	assert.ErrorContains(t, err, "test_error", "Expected error to contain 'test_error', got '%s'", err.Error())
}

func TestUpdateOk(t *testing.T) {
	s, accountRepo, _ := setupAccountsService()

	user := &User{Username: "admin", IsAdmin: true}
	params := &UpdateAccountParams{Username: "account1", Enabled: false}
	accountRepo.On("GetAccountByUsername", mock.Anything, "account1").Return(db.Account{}, nil)
	accountRepo.On("UpdateAccountStatusByUsername", mock.Anything, mock.Anything).Return(
		func(ctx context.Context, params db.UpdateAccountStatusByUsernameParams) error {
			assert.Equal(t, false, params.Enabled, "Expected enabled = false, got '%s'", params.Enabled)
			assert.Equal(t, "account1", params.Username, "Expected username = 'account1', got '%s'", params.Username)
			return nil
		},
	)

	resp, err := s.Update(params, user)

	assert.NoError(t, err, "Expected no error, got '%s'", err)
	assert.NotNil(t, resp, "Expected response, got nil")
	assert.Equal(t, false, resp.Enabled, "Expected enabled = false, got '%s'", resp.Enabled)
}
