package core

import (
	"context"
	"database/sql"

	"log/slog"

	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
)

const (
	ADMIN_USERNAME = "admin"
)

type AuthenticationService struct {
	db.AccountRepository
	*util.Config
}

func (service *AuthenticationService) Authenticate(username string, password string) (token string, err error) {
	var validatePassword func(password string) bool

	// User/Password verification
	if username == ADMIN_USERNAME {
		validatePassword = func(password string) bool {
			return password == service.Config.AdminPassword
		}
	} else {
		account, err := service.AccountRepository.GetAccountByUsername(context.Background(), username)
		if err != nil {
			if err == sql.ErrNoRows {
				slog.Error("Account not found", "Account", username)
				return token, CoreError{
					Code:    ERR_ACCOUNT_NOT_FOUND,
					Message: "account not found",
				}

			}
			slog.Error("Unexpected error while retrieving account by username",
				"Username", username, "Details", err)
			return token, err
		}
		hashedPassword := account.HashedPassword
		validatePassword = func(password string) bool {
			return util.VerifyPassword(hashedPassword, password)
		}
	}

	if !validatePassword(password) {
		slog.Error("Given password does not match with the stored password")
		return token, CoreError{
			Code:    ERR_UNAUTHORIZED,
			Message: "invalid credentials",
		}
	}

	// Generate a JWT token
	token, err = util.GenerateJWT(username, service.Config.TokenSymmetricKey)
	if err != nil {
		slog.Error("Error generating JWT token", "Details", err)
		return token, err
	}

	slog.Info("Login successful", "Username", username)
	return token, nil
}

func (service *AuthenticationService) Validate(token string) (*User, error) {
	username, err := util.ValidateJWT(token, service.Config.TokenSymmetricKey)
	if err != nil {
		slog.Error("Error validating JWT Token", "Details", err)
		return (*User)(nil), nil
	}
	return &User{
		Username: username,
		IsAdmin:  username == ADMIN_USERNAME,
	}, err
}
