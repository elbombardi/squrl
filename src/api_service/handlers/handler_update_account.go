package handlers

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/elbombardi/squrl/src/api_service/api/operations"
	"github.com/elbombardi/squrl/src/db"
	"github.com/go-openapi/runtime/middleware"
)

func (h *Handlers) HandleUpdateAccount(params operations.PutAccountParams) middleware.Responder {
	//Validate params
	err := validateUpdateAccountParams(params)
	if err != nil {
		return operations.NewPutAccountBadRequest().WithPayload(&operations.PutAccountBadRequestBody{
			Error: err.Error()})
	}

	//Check if the Admin API key is valid
	if params.XAPIKEY != h.Config.AdminPassword {
		return operations.NewPostAccountUnauthorized().WithPayload(&operations.PostAccountUnauthorizedBody{
			Error: "invalid x-api-key header"})
	}

	//Check if the customer exists
	_, err = h.AccountRepository.GetAccountByUsername(context.Background(), *params.Body.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return operations.NewPutAccountNotFound().WithPayload(&operations.PutAccountNotFoundBody{
				Error: "customer not found for this username: " + *params.Body.Username})
		}
		return internalErrorInUpdateAccount(err)
	}

	//Update customer
	err = h.AccountRepository.UpdateAccountStatusByUsername(context.Background(), db.UpdateAccountStatusByUsernameParams{
		Username: *params.Body.Username,
		Enabled:  encodeStatus(*params.Body.Status),
	})
	if err != nil {
		return internalErrorInUpdateAccount(err)
	}
	return operations.NewPutAccountOK().WithPayload(&operations.PutAccountOKBody{
		Status: *params.Body.Status,
	})
}

func validateUpdateAccountParams(params operations.PutAccountParams) error {
	if params.Body.Username == nil {
		return errors.New("missing username")
	}
	if params.Body.Status == nil {
		return errors.New("missing status")
	}
	if *params.Body.Status != "active" && *params.Body.Status != "inactive" {
		return errors.New("invalid status, should be one of the two values: 'active', 'inactive'")
	}
	if params.XAPIKEY == "" {
		return errors.New("missing x-api-key header")
	}
	return nil
}

func internalErrorInUpdateAccount(err error) middleware.Responder {
	log.Println("Error updating customer: ", err)
	return operations.NewPutAccountInternalServerError().WithPayload(&operations.PutAccountInternalServerErrorBody{
		Error: err.Error()})
}
