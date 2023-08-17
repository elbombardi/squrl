package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/elbombardi/squrl/src/api_service/api/operations"
	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
	"github.com/go-openapi/runtime/middleware"
)

func (h *Handlers) HandleCreateURL(params operations.PostShortURLParams) middleware.Responder {
	//Validate params
	err := validateCreateShortURLParams(params)
	if err != nil {
		return operations.NewPostShortURLBadRequest().WithPayload(&operations.PostShortURLBadRequestBody{
			Error: err.Error()})
	}

	//Check if the account API key is valid
	account, err := h.AccountRepository.GetAccountByApiKey(context.Background(), params.XAPIKEY)
	if err != nil {
		if err == sql.ErrNoRows {
			return operations.NewPostShortURLUnauthorized().WithPayload(&operations.PostShortURLUnauthorizedBody{
				Error: "Invalid API Key"})
		}
		return internalErrorInCreateShortURL(err)
	}

	//Check if the account is active
	if !account.Enabled {
		return operations.NewPostShortURLUnauthorized().WithPayload(&operations.PostShortURLUnauthorizedBody{
			Error: "Customer disabled"})
	}

	//Generate short URL key
	shortURLKey, err := h.generateShortURLKey()
	if err != nil {
		return internalErrorInCreateShortURL(err)
	}

	//Insert the new short URL in the database
	err = h.URLRepository.InsertNewURL(context.Background(), db.InsertNewURLParams{
		ShortUrlKey: sql.NullString{String: shortURLKey, Valid: true},
		LongUrl:     *params.Body.LongURL,
		AccountID:   account.ID,
	})
	if err != nil {
		return internalErrorInCreateShortURL(err)
	}

	//Return the short URL
	return operations.NewPostShortURLOK().WithPayload(&operations.PostShortURLOKBody{
		ShortURL:    fmt.Sprintf("%v/%v/%v", h.Config.RedirectionBaseURL, account.Prefix, shortURLKey),
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
	err := util.ValidateURL(*params.Body.LongURL)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handlers) generateShortURLKey() (string, error) {
	apiKey := util.GenerateRandomString(6)
	//Check if the url key is unique
	exists, err := h.AccountRepository.CheckApiKeyExists(context.Background(), apiKey)
	if err != nil {
		return "", err
	}
	if exists {
		return h.generateShortURLKey()
	}
	return apiKey, nil
}

func internalErrorInCreateShortURL(err error) middleware.Responder {
	log.Println("Error creating short URL: ", err)
	return operations.NewPostShortURLInternalServerError().WithPayload(&operations.PostShortURLInternalServerErrorBody{
		Error: err.Error()})
}
