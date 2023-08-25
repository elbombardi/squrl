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
func (h *Handlers) HandleLogin(params general.LoginParams) middleware.Responder {

	if params.Login == nil {
		return general.NewLoginBadRequest().WithPayload(&models.Error{
			Message: "Request body is required",
		})
	}

	token, err := h.Authenticate(params.Login.Username, params.Login.Password)

	if err != nil {
		coreError, ok := err.(core.CoreError)
		switch {
		case ok && coreError.Code == core.ErrBadParams:
			return general.NewLoginBadRequest().WithPayload(&models.Error{
				Message: coreError.Message})
		case ok && coreError.Code == core.ErrAccountNotFound:
			fallthrough
		case ok && coreError.Code == core.ErrAccountDisabled:
			fallthrough
		case ok && coreError.Code == core.ErrUnauthorized:
			return general.NewLoginUnauthorized().WithPayload(&models.Error{
				Message: "Invalid credentials"})
		default:
			return general.NewLoginInternalServerError().WithPayload(&models.Error{
				Message: "Internal server error",
			})
		}
	}

	return general.NewLoginOK().WithPayload(&models.LoginSuccess{
		Success: true,
		Token:   token,
	})
}
