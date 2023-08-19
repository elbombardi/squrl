package handlers

import (
	"context"
	"errors"
	"log/slog"

	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/api/operations/accounts"
	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
	"github.com/go-openapi/runtime/middleware"
)

/*
Handler for the POST /accounts endpoint
*/
func (handlers *Handlers) HandleCreateAccount(params accounts.CreateAccountParams, principal any) middleware.Responder {
	//Check if the request is valid
	err := validateCreateAccountParams(params)
	if err != nil {
		slog.Error("Invalid parameters in create account request", "Details", err)
		return accounts.NewCreateAccountBadRequest().WithPayload(getError(err))
	}
	// Check if the user is authenticated
	if principal == nil {
		return accounts.NewCreateAccountUnauthorized().WithPayload(&models.Error{
			Error: "Unauthorized"})
	}

	//This endpoint is only accessible by the admin
	if principal.(string) != "admin" {
		slog.Error("Unauthorized attempt to access CreateAccount by a non admin user", "Account", principal.(string))
		return accounts.NewCreateAccountUnauthorized().WithPayload(&models.Error{
			Error: "Unauthorized"})
	}

	//Check if the username is unique
	exists, err := handlers.AccountRepository.CheckUsernameExists(context.Background(), *params.Account.Username)
	if err != nil {
		slog.Error("Unexpected error while checking if username exists", "Details", err)
		return internalErrorInCreateAccount()
	}
	if exists {
		slog.Error("Username already exists", "Username", *params.Account.Username)
		return accounts.NewCreateAccountBadRequest().WithPayload(getError(err))
	}

	// Generate a prefix for the account
	prefix, err := handlers.generatePrefix()
	if err != nil {
		slog.Error("Unexpected error while generating prefix for the new account", "Details", err)
		return internalErrorInCreateAccount()
	}

	// Generate a password
	password, hashedPassword := handlers.generatePassword()
	if err != nil {
		slog.Error("Unexpected error while generating password", "Details", err)
		return internalErrorInCreateAccount()
	}

	// Insert the new account
	err = handlers.AccountRepository.InsertNewAccount(context.Background(), db.InsertNewAccountParams{
		Prefix:         prefix,
		HashedPassword: hashedPassword,
		Username:       *params.Account.Username,
		Email:          *params.Account.Email,
	})
	if err != nil {
		slog.Error("Unexpected error while inserting new account in DB", "Details", err)
		return internalErrorInCreateAccount()
	}

	// Return response
	slog.Info("New account created successfully", "Username", *params.Account.Username)
	return accounts.NewCreateAccountOK().WithPayload(&models.AccountCreated{
		Password: password,
		Prefix:   prefix,
	})
}

func validateCreateAccountParams(params accounts.CreateAccountParams) error {
	if params.Account.Username == nil {
		return errors.New("missing username")
	}
	err := util.ValidateUsername(*params.Account.Username)
	if err != nil {
		return err
	}
	if params.Account.Email == nil {
		return errors.New("missing email")
	}
	err = util.ValidateEmail(*params.Account.Email)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handlers) generatePrefix() (string, error) {
	prefix := util.GenerateRandomString(3)
	//Check if the prefix is unique
	exists, err := h.AccountRepository.CheckPrefixExists(context.Background(), prefix)
	if err != nil {
		return "", err
	}
	if exists {
		return h.generatePrefix()
	}
	return prefix, nil
}

func (h *Handlers) generatePassword() (string, string) {
	password := util.GenerateRandomString(20)
	hashedPassword := util.HashPassword(password)
	return password, hashedPassword
}

func internalErrorInCreateAccount() middleware.Responder {
	return accounts.NewCreateAccountInternalServerError().WithPayload(&models.Error{
		Error: "Unexpected server internal error",
	})
}
