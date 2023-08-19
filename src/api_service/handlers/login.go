package handlers

import (
	"context"
	"database/sql"
	"log/slog"

	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/api/operations/general"
	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/go-openapi/runtime/middleware"
)

/*
Handler for POST /login endpoint.

Authentication using username/password

Returns a JWT token
*/
func (handlers *Handlers) HandleLogin(login general.LoginParams) middleware.Responder {
	var validatePassword func(givenPassword string) bool

	// User/Password verification
	if *login.Login.Username == "admin" {
		validatePassword = func(givenPassword string) bool {
			return givenPassword == handlers.Config.AdminPassword
		}
	} else {
		account, err := handlers.AccountRepository.GetAccountByUsername(context.Background(), *login.Login.Username)
		if err != nil {
			if err == sql.ErrNoRows {
				slog.Error("Account not found", "Account", *login.Login.Username)
				return general.NewLoginUnauthorized().WithPayload(&models.Error{
					Error: "Invalid username or password",
				})
			}
			slog.Error("Unexpected error while retrieving account by username", "Username", *login.Login.Username, "Details", err)
			return general.NewLoginInternalServerError().WithPayload(&models.Error{
				Error: "Internal server error",
			})
		}
		hashedPassword := account.HashedPassword
		validatePassword = func(givenPassword string) bool {
			return util.VerifyPassword(hashedPassword, givenPassword)
		}
	}
	if !validatePassword(*login.Login.Password) {
		slog.Error("Given password does not match with the stored password")
		return general.NewLoginUnauthorized().WithPayload(&models.Error{
			Error: "Invalid username or password",
		})
	}

	// Generate a JWT token
	token, err := util.GenerateJWT(*login.Login.Username, handlers.Config.TokenSymmetricKey)
	if err != nil {
		slog.Error("Error generating JWT token", "Details", err)
		return general.NewLoginInternalServerError().WithPayload(&models.Error{
			Error: "Internal server error",
		})
	}

	slog.Info("Login successful", "Username", *login.Login.Username)
	return general.NewLoginOK().WithPayload(&models.LoginSuccess{
		Success: true,
		Token:   token,
	})
}
