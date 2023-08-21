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
	*slog.Logger
}

func (s *AuthenticationService) Authenticate(username string, password string) (token string, err error) {
	var validatePassword func(password string) bool

	// User/Password verification
	if username == ADMIN_USERNAME {
		validatePassword = func(password string) bool {
			return password == s.AdminPassword
		}
	} else {
		account, err := s.GetAccountByUsername(context.Background(), username)
		if err != nil {
			if err == sql.ErrNoRows {
				s.Error("Account not found", "Account", username)
				return token, CoreError{
					Code:    ErrAccountNotFound,
					Message: "account not found",
				}
			}
			s.Error("Unexpected error while retrieving account by username",
				"Username", username, "Details", err)
			return token, err
		}
		validatePassword = func(password string) bool {
			return util.VerifyPassword(account.HashedPassword, password)
		}
	}

	if !validatePassword(password) {
		s.Error("Given password does not match with the stored password", "Username", username)
		return token, CoreError{
			Code:    ErrUnauthorized,
			Message: "invalid credentials",
		}
	}

	// Generate a JWT token
	token, err = util.GenerateJWT(username, s.TokenSymmetricKey)
	if err != nil {
		s.Error("Error generating JWT token", "Details", err)
		return token, err
	}

	s.Info("Login successful", "Username", username)
	return token, nil
}

func (s *AuthenticationService) Validate(token string) (*User, error) {
	username, err := util.ValidateJWT(token, s.TokenSymmetricKey)
	if err != nil {
		s.Error("Error validating JWT Token", "Details", err)
		return (*User)(nil), nil
	}
	return &User{
		Username: username,
		IsAdmin:  username == ADMIN_USERNAME,
	}, err
}
