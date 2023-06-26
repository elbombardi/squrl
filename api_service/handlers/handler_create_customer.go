package handlers

import (
	"context"
	"errors"
	"log"

	"github.com/elbombardi/squrl/api_service/api/operations"
	db "github.com/elbombardi/squrl/db/sqlc"
	"github.com/elbombardi/squrl/util"
	"github.com/go-openapi/runtime/middleware"
)

func (h *Handlers) CreateCustomerHandler(params operations.PostCustomerParams) middleware.Responder {
	//Check if the request is valid
	err := validateCreateCustomerParams(params)
	if err != nil {
		return operations.NewPostCustomerBadRequest().WithPayload(&operations.PostCustomerBadRequestBody{
			Error: err.Error()})
	}

	//Check if the Admin API key is valid
	if params.XAPIKEY != *util.ConfigAdminAPIKey() {
		return operations.NewPostCustomerUnauthorized().WithPayload(&operations.PostCustomerUnauthorizedBody{
			Error: "invalid x-api-key header"})
	}

	//Check if the username is unique
	exists, err := h.CustomersRepository.CheckUsernameExists(context.Background(), *params.Customer.Username)
	if err != nil {
		return internalErrorInCreateCustomer(err)
	}
	if exists {
		return operations.NewPostCustomerBadRequest().WithPayload(&operations.PostCustomerBadRequestBody{
			Error: "username already exists"})
	}

	// Generate a prefix
	prefix, err := h.generatePrefix()
	if err != nil {
		return internalErrorInCreateCustomer(err)
	}

	// Generate an API key
	apiKey, err := h.generateAPIKey()
	if err != nil {
		return internalErrorInCreateCustomer(err)
	}

	// Insert the new customer
	err = h.CustomersRepository.InsertNewCustomer(context.Background(), db.InsertNewCustomerParams{
		Prefix:   prefix,
		ApiKey:   apiKey,
		Username: *params.Customer.Username,
		Email:    *params.Customer.Email,
	})
	if err != nil {
		return operations.NewPostCustomerInternalServerError().WithPayload(&operations.PostCustomerInternalServerErrorBody{
			Error: err.Error()})
	}

	// Return response
	return operations.NewPostCustomerOK().WithPayload(&operations.PostCustomerOKBody{
		APIKey: apiKey,
		Prefix: prefix,
	})
}

func validateCreateCustomerParams(params operations.PostCustomerParams) error {
	if params.Customer.Username == nil {
		return errors.New("missing username")
	}
	if params.Customer.Email == nil {
		return errors.New("missing email")
	}
	if params.XAPIKEY == "" {
		return errors.New("missing x-api-key header")
	}
	return nil
}

func (h *Handlers) generatePrefix() (string, error) {
	prefix := util.GenerateRandomString(3)
	//Check if the prefix is unique
	exists, err := h.CustomersRepository.CheckPrefixExists(context.Background(), prefix)
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
	exists, err := h.CustomersRepository.CheckApiKeyExists(context.Background(), apiKey)
	if err != nil {
		return "", err
	}
	if exists {
		return h.generateAPIKey()
	}
	return apiKey, nil
}

func internalErrorInCreateCustomer(err error) middleware.Responder {
	log.Println("Error creating customer: ", err)
	return operations.NewPostCustomerInternalServerError().WithPayload(&operations.PostCustomerInternalServerErrorBody{
		Error: err.Error()})
}
