package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/api/operations/urls"
	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
	"github.com/go-openapi/runtime/middleware"
)

func (handlers *Handlers) HandleCreateURL(params urls.CreateURLParams, principal any) middleware.Responder {
	//Validate params
	err := validateCreateURLParams(params)
	if err != nil {
		return urls.NewCreateURLBadRequest().WithPayload(getError(err))
	}

	//Check if the account API key is valid
	//TODO
	account, err := handlers.AccountRepository.GetAccountByApiKey(context.Background(), principal.(string))
	if err != nil {
		if err == sql.ErrNoRows {
			return urls.NewCreateURLUnauthorized().WithPayload(&models.Error{
				Error: "Invalid API Key"})
		}
		return internalErrorInCreateShortURL(err)
	}

	//Check if the account is active
	if !account.Enabled {
		return urls.NewCreateURLUnauthorized().WithPayload(&models.Error{
			Error: "Customer disabled"})
	}

	//Generate short URL key
	shortURLKey, err := handlers.generateShortURLKey()
	if err != nil {
		return internalErrorInCreateShortURL(err)
	}

	//Insert the new short URL in the database
	err = handlers.URLRepository.InsertNewURL(context.Background(), db.InsertNewURLParams{
		ShortUrlKey: sql.NullString{String: shortURLKey, Valid: true},
		LongUrl:     *params.Body.LongURL,
		AccountID:   account.ID,
	})
	if err != nil {
		return internalErrorInCreateShortURL(err)
	}

	//Return the short URL
	return urls.NewCreateURLOK().WithPayload(&models.URLCreated{
		ShortURL:    fmt.Sprintf("%v/%v/%v", handlers.Config.RedirectionBaseURL, account.Prefix, shortURLKey),
		ShortURLKey: shortURLKey,
	})
}

func validateCreateURLParams(params urls.CreateURLParams) error {
	// API key is required
	//TODO
	// if params.XAPIKEY == "" {
	// 	return errors.New("missing x-api-key header")
	// }
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
	return urls.NewCreateURLInternalServerError().WithPayload(&models.Error{
		Error: err.Error()})
}
