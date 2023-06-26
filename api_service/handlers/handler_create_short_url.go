package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/elbombardi/squrl/api_service/api/operations"
	db "github.com/elbombardi/squrl/db/sqlc"
	"github.com/elbombardi/squrl/util"
	"github.com/go-openapi/runtime/middleware"
)

func (h *Handlers) CreateShortURLHandler(params operations.PostShortURLParams) middleware.Responder {
	//Validate params
	err := validateCreateShortURLParams(params)
	if err != nil {
		return operations.NewPostShortURLBadRequest().WithPayload(&operations.PostShortURLBadRequestBody{
			Error: err.Error()})
	}

	//Check if the customer API key is valid
	customer, err := h.CustomersRepository.GetCustomerByApiKey(context.Background(), params.XAPIKEY)
	if err != nil {
		if err == sql.ErrNoRows {
			return operations.NewPostShortURLUnauthorized().WithPayload(&operations.PostShortURLUnauthorizedBody{
				Error: "Invalid API Key: " + params.XAPIKEY})
		}
		return internalErrorInCreateShortURL(err)
	}

	//Check if the customer is active
	if customer.Status != "e" {
		return operations.NewPostShortURLUnauthorized().WithPayload(&operations.PostShortURLUnauthorizedBody{
			Error: "Customer disabled"})
	}

	//Generate short URL key
	shortURLKey, err := h.generateShortURLKey()
	if err != nil {
		return internalErrorInCreateShortURL(err)
	}

	//Insert the new short URL in the database
	err = h.ShortURLsRepository.InsertNewShortURL(context.Background(), db.InsertNewShortURLParams{
		ShortUrlKey: sql.NullString{String: shortURLKey, Valid: true},
		LongUrl:     *params.Body.LongURL,
		CustomerID:  customer.ID,
	})
	if err != nil {
		return internalErrorInCreateShortURL(err)
	}

	//Return the short URL
	return operations.NewPostShortURLOK().WithPayload(&operations.PostShortURLOKBody{
		ShortURL:    fmt.Sprintf("%v/%v/%v", *util.ConfigRedirectionServerBaseURL(), customer.Prefix, shortURLKey),
		ShortURLKey: shortURLKey,
	})
}

func validateCreateShortURLParams(params operations.PostShortURLParams) error {
	// API key is required
	if params.XAPIKEY == "" {
		return errors.New("missing x-api-key header")
	}
	// URL is required
	if params.Body.LongURL == nil {
		return errors.New("missing param : long_url")
	}
	// URL should be a valid URL
	if valid, err := util.IsValidURL(*params.Body.LongURL); !valid {
		return err
	}
	return nil
}

func (h *Handlers) generateShortURLKey() (string, error) {
	apiKey := util.GenerateRandomString(6)
	//Check if the url key is unique
	exists, err := h.CustomersRepository.CheckApiKeyExists(context.Background(), apiKey)
	if err != nil {
		return "", err
	}
	if exists {
		return h.generateShortURLKey()
	}
	return apiKey, nil
}

func internalErrorInCreateShortURL(err error) middleware.Responder {
	return operations.NewPostShortURLInternalServerError().WithPayload(&operations.PostShortURLInternalServerErrorBody{
		Error: err.Error()})
}
