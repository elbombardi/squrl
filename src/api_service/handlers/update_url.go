package handlers

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/api/operations/urls"
	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
	"github.com/go-openapi/runtime/middleware"
)

/*
Hanlder for the PUT /urls endpoint
*/
func (handlers *Handlers) HandleUpdateShortURL(params urls.UpdateURLParams, principal any) middleware.Responder {
	//Validate params
	err := validateParams(params)
	if err != nil {
		slog.Error("Bad UpdateURL params", "Details", err)
		return urls.NewUpdateURLBadRequest().WithPayload(getError(err))
	}
	if principal == nil {
		return urls.NewUpdateURLUnauthorized().WithPayload(&models.Error{
			Error: "Unauthorized"})
	}

	//Check if the account exists
	account, err := handlers.AccountRepository.GetAccountByUsername(context.Background(), principal.(string))
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("Account not found", "Account", principal.(string))
			return urls.NewUpdateURLUnauthorized().WithPayload(&models.Error{
				Error: "Unauthorized"})
		}
		slog.Error("Unexpected error while retrieving account by username",
			"Account", principal, "Details", err)
		return internalErrorInUpdateShortURL()
	}

	//Check if the customer is active
	if !account.Enabled {
		slog.Info("Account disabled", "Account", principal)
		return urls.NewUpdateURLUnauthorized().WithPayload(&models.Error{
			Error: "Account disabled"})
	}

	url, err := handlers.URLRepository.GetURLByAccountIDAndShortURLKey(context.Background(),
		db.GetURLByAccountIDAndShortURLKeyParams{
			AccountID: account.ID,
			ShortUrlKey: sql.NullString{
				String: *params.Body.ShortURLKey,
				Valid:  true,
			},
		})
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("URL not found", "Account", principal.(string), "ShortURLKey", *params.Body.ShortURLKey)
			return urls.NewUpdateURLNotFound().WithPayload(&models.Error{
				Error: "URL Not Found"})
		}
		slog.Error("Unxpected error while retrieving short URL", "Account", principal.(string),
			"ShortURLKey", *params.Body.ShortURLKey, "Details", err)
		return internalErrorInUpdateShortURL()
	}
	if params.Body.NewLongURL != "" {
		url.LongUrl = params.Body.NewLongURL
		err = handlers.URLRepository.UpdateLongURL(context.Background(), db.UpdateLongURLParams{
			LongUrl: params.Body.NewLongURL,
			ID:      url.ID,
		})
		if err != nil {
			slog.Error("Unexpected error while updating long URL", "Details", err)
			return internalErrorInUpdateShortURL()
		}
	}
	if params.Body.Status != "" {
		url.Enabled = encodeStatus(params.Body.Status)
		err = handlers.URLRepository.UpdateURLStatus(context.Background(), db.UpdateURLStatusParams{
			Enabled: url.Enabled,
			ID:      url.ID,
		})
		if err != nil {
			slog.Error("Unexpected error while updating URL enabled status", "Details", err)
			return internalErrorInUpdateShortURL()
		}
	}
	if params.Body.TrackingStatus != "" {
		url.TrackingEnabled = encodeStatus(params.Body.TrackingStatus)
		err = handlers.URLRepository.UpdateURLTrackingStatus(context.Background(),
			db.UpdateURLTrackingStatusParams{
				TrackingEnabled: url.TrackingEnabled,
				ID:              url.ID,
			})
		if err != nil {
			slog.Error("Unexpected error while updating URL's tracking enabled status: ", "Details", err)
			return internalErrorInUpdateShortURL()
		}
	}
	slog.Info("URL updated successfully", "Account", principal, "Params", *params.Body)
	return urls.NewUpdateURLOK().WithPayload(&urls.UpdateURLOKBody{
		LongURL:        url.LongUrl,
		Status:         decodeStatus(url.Enabled),
		TrackingStatus: decodeStatus(url.TrackingEnabled),
	})
}

func validateParams(params urls.UpdateURLParams) error {
	if params.Authorization == "" {
		return errors.New("missing jwt header")
	}
	if params.Body.ShortURLKey == nil {
		return errors.New("missing parameter : 'short_url_key'")
	}
	if params.Body.NewLongURL != "" {
		err := util.ValidateURL(params.Body.NewLongURL)
		if err != nil {
			return err
		}
	}
	if params.Body.Status != "" &&
		params.Body.Status != "active" &&
		params.Body.Status != "inactive" {
		return errors.New("invalid status, should be one of the two values: 'active', 'inactive'")
	}
	if params.Body.TrackingStatus != "" &&
		params.Body.TrackingStatus != "active" &&
		params.Body.TrackingStatus != "inactive" {
		return errors.New("invalid tracking status, should be one of the two values: 'active', 'inactive'")
	}
	return nil
}

func internalErrorInUpdateShortURL() middleware.Responder {
	return urls.NewUpdateURLInternalServerError().WithPayload(&models.Error{
		Error: "Internal server error",
	})
}
