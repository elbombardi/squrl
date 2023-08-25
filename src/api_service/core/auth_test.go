package core

import (
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupAuthService() (
	*AuthenticationService,
	*db.MockAccountRepository,
	*util.Config,
) {
	accountRepo := new(db.MockAccountRepository)
	config := util.MockConfig()
	return &AuthenticationService{
			AccountRepository: accountRepo,
			Config:            config,
			Logger:            util.NewLogger(config),
		},
		accountRepo,
		config
}

func generateJWT(
	username string,
	tokenSymmetricKey string,
	iss string,
	method jwt.SigningMethod,
	exp int64,
) (string, error) {
	token := jwt.New(method)
	claims := token.Claims.(jwt.MapClaims)
	claims["user"] = username
	claims["iss"] = iss
	claims["exp"] = exp
	tokenString, err := token.SignedString([]byte(tokenSymmetricKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func TestAuthenticateWithAdminAndInvalidePassword(t *testing.T) {
	authService, _, _ := setupAuthService()

	token, err := authService.Authenticate(ADMIN_USERNAME, "invalid")

	assert.Error(t, err, "Expected error, got nil")
	assert.Equal(t, "", token, "Expected empty token, got %s", token)
	assert.Equal(t, ErrUnauthorized, err.(CoreError).Code, "Expected error code %d, got %d", ErrUnauthorized, err.(CoreError).Code)
}

func TestAuthenticateAccountNotFound(t *testing.T) {
	authService, accountRepo, _ := setupAuthService()

	accountRepo.On("GetAccountByUsername", mock.Anything, "unknown").Return(db.Account{}, sql.ErrNoRows)

	token, err := authService.Authenticate("unknown", "password")

	assert.Error(t, err, "Expected error, got nil")
	assert.Equal(t, "", token, "Expected empty token, got %s", token)
	assert.Equal(t, ErrAccountNotFound, err.(CoreError).Code, "Expected error code %d, got %d", ErrAccountNotFound, err.(CoreError).Code)
}

func TestAuthenticateUnexpectedErrorWhileLoadingAccount(t *testing.T) {
	authService, accountRepo, _ := setupAuthService()

	accountRepo.On("GetAccountByUsername", mock.Anything, "unknown").Return(db.Account{}, errors.New("unexpected error"))

	token, err := authService.Authenticate("unknown", "password")

	assert.Error(t, err, "Expected error, got nil")
	assert.Equal(t, "", token, "Expected empty token, got %s", token)
	assert.ErrorContains(t, err, "unexpected error", "Expected error message to contain 'unexpected error', got %s", err)
}

func TestAuthenticateNonAdminInvalidCredentials(t *testing.T) {
	authService, accountRepo, _ := setupAuthService()

	accountRepo.On("GetAccountByUsername", mock.Anything, "account1").Return(db.Account{
		HashedPassword: "wrongpassword!",
		Username:       "account1",
		Enabled:        true,
	}, nil)

	token, err := authService.Authenticate("account1", "password")

	assert.Error(t, err, "Expected error, got nil")
	assert.Equal(t, "", token, "Expected empty token, got %s", token)
	assert.Equal(t, ErrUnauthorized, err.(CoreError).Code, "Expected error code %d, got %d", ErrUnauthorized, err.(CoreError).Code)
}

func TestAuthenticateOk(t *testing.T) {
	authService, accountRepo, config := setupAuthService()

	accountRepo.On("GetAccountByUsername", mock.Anything, "account1").Return(db.Account{
		HashedPassword: util.HashPassword("password"),
		Username:       "account1",
		Enabled:        true,
	}, nil)

	token, err := authService.Authenticate("account1", "password")

	assert.NoError(t, err, "Expected no error, got %s", err)
	assert.NotEqual(t, "", token, "Expected non-empty token, got %s", token)

	fmt.Println(token)

	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.TokenSymmetricKey), nil
	})
	assert.NoError(t, err, "Expected no error, got %s", err)
	assert.True(t, parsed.Valid, "Expected token to be valid, got invalid")
}

func TestValidateWithoutBearerPrefix(t *testing.T) {
	authService, _, _ := setupAuthService()

	user, _ := authService.Validate("invalid")

	assert.Equal(t, (*User)(nil), user, "Expected nil user, got %s", user)
}

func TestValidateInvalidNumberOfSegments(t *testing.T) {
	authService, _, _ := setupAuthService()

	user, _ := authService.Validate("Bearer invalid")

	assert.Equal(t, (*User)(nil), user, "Expected nil user, got %s", user)
}

func TestValidateInvalidSignMethod(t *testing.T) {
	authService, _, config := setupAuthService()

	jwtToken, err := generateJWT("account1", config.TokenSymmetricKey, "squrl",
		jwt.SigningMethodHS512, time.Now().Add(time.Hour*8).Unix())

	fmt.Println(jwtToken, err)
	user, _ := authService.Validate("Bearer " + jwtToken)

	assert.Equal(t, (*User)(nil), user, "Expected nil user, got %s", user)

}

func TestValidateInvalidIss(t *testing.T) {
	authService, _, config := setupAuthService()

	jwtToken, _ := generateJWT("account1", config.TokenSymmetricKey, "badiss",
		jwt.SigningMethodHS256, time.Now().Add(time.Hour*8).Unix())

	user, _ := authService.Validate("Bearer " + jwtToken)

	assert.Equal(t, (*User)(nil), user, "Expected nil user, got %s", user)
}

func TestValidateExpriedToken(t *testing.T) {
	authService, _, config := setupAuthService()

	jwtToken, _ := generateJWT("account1", config.TokenSymmetricKey, "squrl",
		jwt.SigningMethodHS256, time.Now().Add(-time.Hour*8).Unix())

	user, _ := authService.Validate("Bearer " + jwtToken)

	assert.Equal(t, (*User)(nil), user, "Expected nil user, got %s", user)
}

func TestValidateInvalidSecret(t *testing.T) {
	authService, _, _ := setupAuthService()

	jwtToken, _ := generateJWT("account1", "somerandomkey", "squrl",
		jwt.SigningMethodHS256, time.Now().Add(time.Hour*8).Unix())

	user, _ := authService.Validate("Bearer " + jwtToken)

	assert.Equal(t, (*User)(nil), user, "Expected nil user, got %s", user)
}

func TestValidateOk(t *testing.T) {
	authService, _, config := setupAuthService()

	jwtToken, _ := generateJWT("account1", config.TokenSymmetricKey, "squrl",
		jwt.SigningMethodHS256, time.Now().Add(time.Hour*8).Unix())

	user, _ := authService.Validate("Bearer " + jwtToken)

	assert.NotEqual(t, (*User)(nil), user, "Expected nil user, got %s", user)
}
