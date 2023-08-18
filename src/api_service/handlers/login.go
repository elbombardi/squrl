package handlers

import (
	"context"
	"database/sql"

	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/api/operations/general"
	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/go-openapi/runtime/middleware"
)

func (handlers *Handlers) HandleLogin(login general.LoginParams) middleware.Responder {
	// Authenicate using the username and password
	var validatePassword func(givenPassword string) bool

	if *login.Login.Username == "admin" {
		validatePassword = func(givenPassword string) bool {
			return givenPassword == handlers.Config.AdminPassword
		}
	} else {
		account, err := handlers.AccountRepository.GetAccountByUsername(context.Background(), *login.Login.Username)
		if err != nil {
			if err == sql.ErrNoRows {
				util.Error("account not found for this username: " + *login.Login.Username)
				return general.NewLoginUnauthorized().WithPayload(&models.Error{
					Error: "account not found for this username: " + *login.Login.Username})
			}
			util.Error("Error getting account by username: ", err)
			return general.NewLoginInternalServerError().WithPayload(&models.Error{
				Error: "Unexpected error getting account by username",
			})
		}
		hashedPassword := account.HashedPassword
		validatePassword = func(givenPassword string) bool {
			return util.VerifyPassword(hashedPassword, givenPassword)
		}
	}
	if !validatePassword(*login.Login.Password) {
		util.Error("invalid username or password")
		return general.NewLoginUnauthorized().WithPayload(&models.Error{
			Error: "invalid username or password",
		})
	}

	// Generate a JWT token
	token, err := util.GenerateJWT(*login.Login.Username, handlers.Config.TokenSymmetricKey)
	if err != nil {
		util.Error("error generating JWT token: ", err)
		return general.NewLoginInternalServerError().WithPayload(&models.Error{
			Error: "error generating JWT token",
		})
	}

	util.Info("Login successful for user: ", *login.Login.Username)
	return general.NewLoginOK().WithPayload(&models.LoginSuccess{
		Success: true,
		Token:   token,
	})
}
