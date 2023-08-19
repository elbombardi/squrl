package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"

	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/api/operations/urls"
	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
	"github.com/go-openapi/runtime/middleware"
)

/*
Handler for the POST /urls endpoint
*/
func (handlers *Handlers) HandleCreateURL(params urls.CreateURLParams, principal any) middleware.Responder {
	//Validate params
	err := validateCreateURLParams(params)
	if err != nil {
		slog.Error("Bad CreateURL request params", "Details", err)
		return urls.NewCreateURLBadRequest().WithPayload(getError(err))
	}
	// Check if the user is authenticated
	if principal == nil {
		return urls.NewCreateURLUnauthorized().WithPayload(&models.Error{
			Error: "Unauthorized"})
	}

	//Check if the account exists
	account, err := handlers.AccountRepository.GetAccountByUsername(context.Background(), principal.(string))
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("Account not found", "Username", principal.(string))
			return urls.NewCreateURLUnauthorized().WithPayload(&models.Error{
				Error: "Account not found for this username: " + principal.(string)})
		}
		slog.Error("Unexpected error while retrieving account by username", "Username", principal.(string), "Detail", err)
		return internalErrorInCreateShortURL()
	}

	//Check if the account is active
	if !account.Enabled {
		slog.Info("Account disabled", "Username", principal.(string))
		return urls.NewCreateURLUnauthorized().WithPayload(&models.Error{
			Error: "Account disabled"})
	}

	//Generate short URL key
	shortURLKey, err := handlers.generateShortURLKey(account)
	if err != nil {
		slog.Error("Unexpected error while generating short URL", "Details", err)
		return internalErrorInCreateShortURL()
	}

	//Insert the new short URL in the database
	err = handlers.URLRepository.InsertNewURL(context.Background(), db.InsertNewURLParams{
		ShortUrlKey: sql.NullString{String: shortURLKey, Valid: true},
		LongUrl:     *params.Body.LongURL,
		AccountID:   account.ID,
	})
	if err != nil {
		slog.Error("Error inserting new URL in DB", "Details", err)
		return internalErrorInCreateShortURL()
	}

	slog.Info("New short URL created successfully", "Account", principal, "Params", *params.Body)
	//Return the short URL
	return urls.NewCreateURLOK().WithPayload(&models.URLCreated{
		ShortURL:    fmt.Sprintf("%v/%v/%v", handlers.Config.RedirectionBaseURL, account.Prefix, shortURLKey),
		ShortURLKey: shortURLKey,
	})
}

func validateCreateURLParams(params urls.CreateURLParams) error {
	if params.Authorization == "" {
		return errors.New("missing jwt header")
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

func (h *Handlers) generateShortURLKey(account db.Account) (string, error) {
	shortURLKey := util.GenerateRandomString(6)
	//Check if the url key is unique
	exists, err := h.URLRepository.CheckShortUrlKeyExists(context.Background(), db.CheckShortUrlKeyExistsParams{
		ShortUrlKey: sql.NullString{String: shortURLKey, Valid: true},
		AccountID:   account.ID,
	})
	if err != nil {
		return "", err
	}
	if exists {
		return h.generateShortURLKey(account)
	}
	return shortURLKey, nil
}

func internalErrorInCreateShortURL() middleware.Responder {
	return urls.NewCreateURLInternalServerError().WithPayload(&models.Error{
		Error: "Unexpected server internal error",
	})
}
