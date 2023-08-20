package handlers

import (
	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/api/operations/accounts"
	"github.com/elbombardi/squrl/src/api_service/core"
	"github.com/go-openapi/runtime/middleware"
)

// Handler for the PUT /accounts endpoint
func (handlers *Handlers) HandleUpdateAccount(params accounts.UpdateAccountParams, principal any) middleware.Responder {

	if params.Body == nil {
		return accounts.NewUpdateAccountBadRequest().WithPayload(&models.Error{
			Message: "Request body is required",
		})
	}

	if params.Body.Status != "active" && params.Body.Status != "inactive" {
		return accounts.NewUpdateAccountBadRequest().WithPayload(&models.Error{
			Message: "Invalid status, should be one of the two values: 'active', 'inactive'",
		})
	}

	response, err := handlers.AccountsManager.Update(&core.UpdateAccountParams{
		Enabled:  encodeStatus(params.Body.Status).Value,
		Username: params.Body.Username,
	}, principal.(*core.User))

	if err != nil {
		coreError, ok := err.(*core.CoreError)
		switch {
		case ok && coreError.Code == core.ERR_BAD_PARAMS:
			return accounts.NewUpdateAccountBadRequest().WithPayload(&models.Error{
				Message: coreError.Message,
			})
		case ok && coreError.Code == core.ERR_UNAUTHORIZED:
			return accounts.NewUpdateAccountUnauthorized().WithPayload(&models.Error{
				Message: "Unauthorized access"})
		case ok && coreError.Code == core.ERR_ACCOUNT_NOT_FOUND:
			return accounts.NewUpdateAccountUnauthorized().WithPayload(&models.Error{
				Message: "Account not found for this username: " + params.Body.Username})
		default:
			return accounts.NewUpdateAccountInternalServerError().WithPayload(&models.Error{
				Message: "Internal server error",
			})
		}
	}

	return accounts.NewUpdateAccountOK().WithPayload(&models.AccountUpdated{
		Status: decodeStatus(response.Enabled),
	})
}
