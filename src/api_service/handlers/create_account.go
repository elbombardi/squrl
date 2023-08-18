package handlers

import (
	"context"
	"errors"
	"log"

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
		return accounts.NewCreateAccountBadRequest().WithPayload(getError(err))
	}

	//Check if the Admin API key is valid
	if params.Authorization != handlers.Config.AdminPassword {
		return accounts.NewCreateAccountUnauthorized().WithPayload(getError(err))
	}

	//Check if the username is unique
	exists, err := handlers.AccountRepository.CheckUsernameExists(context.Background(), *params.Account.Username)
	if err != nil {
		return internalErrorInCreateAccount(err)
	}
	if exists {
		return accounts.NewCreateAccountBadRequest().WithPayload(getError(err))
	}

	// Generate a prefix
	prefix, err := handlers.generatePrefix()
	if err != nil {
		return internalErrorInCreateAccount(err)
	}

	// Generate an API key
	apiKey, err := handlers.generateAPIKey()
	if err != nil {
		return internalErrorInCreateAccount(err)
	}

	// Insert the new account
	err = handlers.AccountRepository.InsertNewAccount(context.Background(), db.InsertNewAccountParams{
		Prefix:   prefix,
		ApiKey:   apiKey,
		Username: *params.Account.Username,
		Email:    *params.Account.Email,
	})
	if err != nil {
		return accounts.NewCreateAccountInternalServerError().WithPayload(getError(err))
	}

	// Return response
	return accounts.NewCreateAccountOK().WithPayload(&models.AccountCreated{
		APIKey: apiKey,
		Prefix: prefix,
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
	if params.Authorization == "" {
		return errors.New("missing jwt header")
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

func (h *Handlers) generateAPIKey() (string, error) {
	apiKey := util.GenerateRandomString(32)
	//Check if the Api key is unique
	exists, err := h.AccountRepository.CheckApiKeyExists(context.Background(), apiKey)
	if err != nil {
		return "", err
	}
	if exists {
		return h.generateAPIKey()
	}
	return apiKey, nil
}

func internalErrorInCreateAccount(err error) middleware.Responder {
	log.Println("Error creating account: ", err)
	return accounts.NewCreateAccountInternalServerError().WithPayload(getError(err))
}
