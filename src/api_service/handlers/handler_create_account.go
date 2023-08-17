package handlers

import (
	"context"
	"errors"
	"log"

	"github.com/elbombardi/squrl/src/api_service/api/operations"
	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
	"github.com/go-openapi/runtime/middleware"
)

func (h *Handlers) HandleCreateAccount(params operations.PostAccountParams) middleware.Responder {
	//Check if the request is valid
	err := validateCreateAccountParams(params)
	if err != nil {
		return operations.NewPostAccountBadRequest().WithPayload(&operations.PostAccountBadRequestBody{
			Error: err.Error()})
	}

	//Check if the Admin API key is valid
	if params.Authorization != h.Config.AdminPassword {
		return operations.NewPostAccountUnauthorized().WithPayload(&operations.PostAccountUnauthorizedBody{
			Error: "invalid jwt header"})
	}

	//Check if the username is unique
	exists, err := h.AccountRepository.CheckUsernameExists(context.Background(), *params.Account.Username)
	if err != nil {
		return internalErrorInCreateAccount(err)
	}
	if exists {
		return operations.NewPostAccountBadRequest().WithPayload(&operations.PostAccountBadRequestBody{
			Error: "username already exists"})
	}

	// Generate a prefix
	prefix, err := h.generatePrefix()
	if err != nil {
		return internalErrorInCreateAccount(err)
	}

	// Generate an API key
	apiKey, err := h.generateAPIKey()
	if err != nil {
		return internalErrorInCreateAccount(err)
	}

	// Insert the new account
	err = h.AccountRepository.InsertNewAccount(context.Background(), db.InsertNewAccountParams{
		Prefix:   prefix,
		ApiKey:   apiKey,
		Username: *params.Account.Username,
		Email:    *params.Account.Email,
	})
	if err != nil {
		return operations.NewPostAccountInternalServerError().WithPayload(&operations.PostAccountInternalServerErrorBody{
			Error: err.Error()})
	}

	// Return response
	return operations.NewPostAccountOK().WithPayload(&operations.PostAccountOKBody{
		APIKey: apiKey,
		Prefix: prefix,
	})
}

func validateCreateAccountParams(params operations.PostAccountParams) error {
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
	return operations.NewPostAccountInternalServerError().WithPayload(&operations.PostAccountInternalServerErrorBody{
		Error: err.Error()})
}
