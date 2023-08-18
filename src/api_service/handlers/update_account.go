package handlers

import (
	"context"
	"database/sql"
	"errors"

	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/api/operations/accounts"
	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
	"github.com/go-openapi/runtime/middleware"
)

func (handlers *Handlers) HandleUpdateAccount(params accounts.UpdateAccountParams, principal any) middleware.Responder {
	//Validate params
	err := validateUpdateAccountParams(params)
	if err != nil {
		util.Error("validating UpdateAccount params: ", err)
		return accounts.NewUpdateAccountBadRequest().WithPayload(getError(err))
	}
	// Check if the user is authenticated
	if principal == nil {
		return accounts.NewUpdateAccountUnauthorized().WithPayload(&models.Error{
			Error: "Unauthorized"})
	}
	//This endpoint is only accessible by the admin
	if principal.(string) != "admin" {
		util.Error("Unauthorized attempt to access UpdateAccount by a non admin user", principal.(string))
		return accounts.NewCreateAccountUnauthorized().WithPayload(&models.Error{
			Error: "Unauthorized"})
	}

	//Check if the customer exists
	_, err = handlers.AccountRepository.GetAccountByUsername(context.Background(), *params.Body.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			util.Error("account not found for this username: ", *params.Body.Username)
			return accounts.NewUpdateAccountNotFound().WithPayload(&models.Error{
				Error: "account not found for this username: " + *params.Body.Username})
		}
		util.Error("getting account by username: ", err)
		return internalErrorInUpdateAccount(err)
	}

	//Update customer
	err = handlers.AccountRepository.UpdateAccountStatusByUsername(context.Background(), db.UpdateAccountStatusByUsernameParams{
		Username: *params.Body.Username,
		Enabled:  encodeStatus(*params.Body.Status),
	})
	if err != nil {
		util.Error("updating account: ", err)
		return internalErrorInUpdateAccount(err)
	}

	util.Info("Account updated successfully: ", *params.Body.Username)
	return accounts.NewUpdateAccountOK().WithPayload("ok")
}

func validateUpdateAccountParams(params accounts.UpdateAccountParams) error {
	if params.Authorization == "" {
		return errors.New("missing jwt header")
	}
	if params.Body.Username == nil {
		return errors.New("missing username")
	}
	if params.Body.Status == nil {
		return errors.New("missing status")
	}
	if *params.Body.Status != "active" && *params.Body.Status != "inactive" {
		return errors.New("invalid status, should be one of the two values: 'active', 'inactive'")
	}
	return nil
}

func internalErrorInUpdateAccount(err error) middleware.Responder {
	return accounts.NewUpdateAccountInternalServerError().WithPayload(getError(err))
}
