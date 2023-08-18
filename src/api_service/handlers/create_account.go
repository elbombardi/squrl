package handlers

import (
	"context"
	"errors"

	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/api/operations/accounts"
	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
	"github.com/go-openapi/runtime/middleware"
)

func (handlers *Handlers) HandleCreateAccount(params accounts.CreateAccountParams, principal any) middleware.Responder {
	//Check if the request is valid
	err := validateCreateAccountParams(params)
	if err != nil {
		util.Error("Invalid parameters in create account request", err)
		return accounts.NewCreateAccountBadRequest().WithPayload(getError(err))
	}
	// Check if the user is authenticated
	if principal == nil {
		return accounts.NewCreateAccountUnauthorized().WithPayload(&models.Error{
			Error: "Unauthorized"})
	}

	//This endpoint is only accessible by the admin
	if principal.(string) != "admin" {
		util.Error("Unauthorized attempt to access CreateAccount by a non admin user", principal.(string))
		return accounts.NewCreateAccountUnauthorized().WithPayload(&models.Error{
			Error: "Unauthorized"})
	}

	//Check if the username is unique
	exists, err := handlers.AccountRepository.CheckUsernameExists(context.Background(), *params.Account.Username)
	if err != nil {
		util.Error("checking if username exists", err)
		return internalErrorInCreateAccount(err)
	}
	if exists {
		util.Error("Username already exists", *params.Account.Username)
		return accounts.NewCreateAccountBadRequest().WithPayload(getError(err))
	}

	// Generate a prefix
	prefix, err := handlers.generatePrefix()
	if err != nil {
		util.Error("Username already exists", *params.Account.Username)
		return internalErrorInCreateAccount(err)
	}

	// Generate an API key
	password, hashedPassword := handlers.generatePassword()
	if err != nil {
		util.Error("Error generating password", err)
		return internalErrorInCreateAccount(err)
	}

	// Insert the new account
	err = handlers.AccountRepository.InsertNewAccount(context.Background(), db.InsertNewAccountParams{
		Prefix:         prefix,
		HashedPassword: hashedPassword,
		Username:       *params.Account.Username,
		Email:          *params.Account.Email,
	})
	if err != nil {
		util.Error("Error inserting new account", err)
		return accounts.NewCreateAccountInternalServerError().WithPayload(getError(err))
	}

	// Return response
	util.Info("New account created successfully: '%s' \n", *params.Account.Username)
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

func internalErrorInCreateAccount(err error) middleware.Responder {
	return accounts.NewCreateAccountInternalServerError().WithPayload(getError(err))
}
