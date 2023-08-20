package handlers

import (
	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/api/operations/general"
	"github.com/elbombardi/squrl/src/api_service/core"
	"github.com/go-openapi/runtime/middleware"
)

// Handler for POST /login endpoint.
// Authentication using username/password
//
// Returns a JWT token
func (handlers *Handlers) HandleLogin(params general.LoginParams) middleware.Responder {

	token, err := handlers.Authenticator.Authenticate(params.Login.Username, params.Login.Password)

	if err != nil {
		coreError, ok := err.(*core.CoreError)
		switch {
		case ok && coreError.Code == core.ERR_ACCOUNT_NOT_FOUND:
			return general.NewLoginBadRequest().WithPayload(&models.Error{
				Message: "Invalid credentials"})
		case ok && coreError.Code == core.ERR_UNAUTHORIZED:
			return general.NewLoginBadRequest().WithPayload(&models.Error{
				Message: "Invalid credentials"})
		default:
			return general.NewLoginInternalServerError().WithPayload(&models.Error{
				Message: "Unexpected server internal error",
			})
		}
	}

	return general.NewLoginOK().WithPayload(&models.LoginSuccess{
		Success: true,
		Token:   token,
	})
}
