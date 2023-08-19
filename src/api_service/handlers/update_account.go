package handlers

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/api/operations/accounts"
	"github.com/elbombardi/squrl/src/db"
	"github.com/go-openapi/runtime/middleware"
)

/*
Handler for the PUT /accounts endpoint
*/
func (handlers *Handlers) HandleUpdateAccount(params accounts.UpdateAccountParams, principal any) middleware.Responder {

	// Validate params
	err := validateUpdateAccountParams(params)
	if err != nil {
		slog.Error("Bad UpdateAccount params", "Details", err)
		return accounts.NewUpdateAccountBadRequest().WithPayload(getError(err))
	}
	// Check if the user is authenticated
	if principal == nil {
		return accounts.NewUpdateAccountUnauthorized().WithPayload(&models.Error{
			Error: "Unauthorized"})
	}
	// This endpoint is only accessible by the admin
	if principal.(string) != "admin" {
		slog.Error("Unauthorized attempt to access UpdateAccount by a non admin user", "User", principal.(string))
		return accounts.NewCreateAccountUnauthorized().WithPayload(&models.Error{
			Error: "Unauthorized"})
	}

	// Check if the customer exists
	_, err = handlers.AccountRepository.GetAccountByUsername(context.Background(), *params.Body.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("Account not found for this username", "Username", *params.Body.Username)
			return accounts.NewUpdateAccountNotFound().WithPayload(&models.Error{
				Error: "Account not found for this username: " + *params.Body.Username})
		}
		slog.Error("Unexpected error while retrieving account by username", "Details", err)
		return internalErrorInUpdateAccount()
	}

	//Update customer
	err = handlers.AccountRepository.UpdateAccountStatusByUsername(context.Background(), db.UpdateAccountStatusByUsernameParams{
		Username: *params.Body.Username,
		Enabled:  encodeStatus(*params.Body.Status),
	})
	if err != nil {
		slog.Error("Unexpected error while updating account", "Details", err)
		return internalErrorInUpdateAccount()
	}

	slog.Info("Account updated successfully", "Account", principal, "Params", *params.Body)
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

func internalErrorInUpdateAccount() middleware.Responder {
	return accounts.NewUpdateAccountInternalServerError().WithPayload(&models.Error{
		Error: "Unexpected server internal error",
	})
}
