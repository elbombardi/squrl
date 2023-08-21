package handlers

import (
	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/api/operations/accounts"
	"github.com/elbombardi/squrl/src/api_service/core"
	"github.com/go-openapi/runtime/middleware"
)

// Handler for the POST /accounts endpoint
func (h *Handlers) HandleCreateAccount(params accounts.CreateAccountParams, principal any) middleware.Responder {

	if params.Account == nil {
		return accounts.NewCreateAccountBadRequest().WithPayload(&models.Error{
			Message: "Account information is required",
		})
	}

	response, err := h.AccountsManager.Create(&core.CreateAccountParams{
		Username: params.Account.Username,
		Email:    params.Account.Email,
	}, principal.(*core.User))

	if err != nil {
		coreErr, ok := err.(core.CoreError)
		switch {
		case ok && coreErr.Code == core.ErrBadParams:
			return accounts.NewCreateAccountBadRequest().WithPayload(&models.Error{
				Message: coreErr.Message,
			})
		case ok && coreErr.Code == core.ErrUnauthorized:
			return accounts.NewCreateAccountUnauthorized().WithPayload(&models.Error{
				Message: "Unauthorized access"})
		default:
			return accounts.NewCreateAccountInternalServerError().WithPayload(&models.Error{
				Message: "Unexpected server internal error",
			})
		}
	}

	return accounts.NewCreateAccountOK().WithPayload(&models.AccountCreated{
		Password: response.Password,
		Prefix:   response.Prefix,
	})
}
