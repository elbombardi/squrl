package handlers

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/api/operations/accounts"
	"github.com/elbombardi/squrl/src/db"
	"github.com/go-openapi/runtime/middleware"
)

func (handlers *Handlers) HandleUpdateAccount(params accounts.UpdateAccountParams, principal any) middleware.Responder {
	//Validate params
	err := validateUpdateAccountParams(params)
	if err != nil {
		return accounts.NewUpdateAccountBadRequest().WithPayload(getError(err))
	}

	//Check if the Admin API key is valid
	//TODO
	// if params.XAPIKEY != h.Config.AdminPassword {
	// 	return accounts.NewUpdateAccountUnauthorized().WithPayload(&accounts.UpdateAccountUnauthorizedBody{
	// 		Error: "invalid x-api-key header"})
	// }

	//Check if the customer exists
	_, err = handlers.AccountRepository.GetAccountByUsername(context.Background(), *params.Body.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return accounts.NewUpdateAccountNotFound().WithPayload(&models.Error{
				Error: "account not found for this username: " + *params.Body.Username})
		}
		return internalErrorInUpdateAccount(err)
	}

	//Update customer
	err = handlers.AccountRepository.UpdateAccountStatusByUsername(context.Background(), db.UpdateAccountStatusByUsernameParams{
		Username: *params.Body.Username,
		Enabled:  encodeStatus(*params.Body.Status),
	})
	if err != nil {
		return internalErrorInUpdateAccount(err)
	}
	return accounts.NewUpdateAccountOK().WithPayload("ok")
}

func validateUpdateAccountParams(params accounts.UpdateAccountParams) error {
	if params.Body.Username == nil {
		return errors.New("missing username")
	}
	if params.Body.Status == nil {
		return errors.New("missing status")
	}
	if *params.Body.Status != "active" && *params.Body.Status != "inactive" {
		return errors.New("invalid status, should be one of the two values: 'active', 'inactive'")
	}
	// TODO
	// if params.XAPIKEY == "" {
	// 	return errors.New("missing x-api-key header")
	// }
	return nil
}

func internalErrorInUpdateAccount(err error) middleware.Responder {
	log.Println("Error updating customer: ", err)
	return accounts.NewUpdateAccountInternalServerError().WithPayload(getError(err))
}
