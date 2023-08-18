package handlers

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/elbombardi/squrl/src/api_service/api/models"
	"github.com/elbombardi/squrl/src/api_service/api/operations/urls"
	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
	"github.com/go-openapi/runtime/middleware"
)

func (handlers *Handlers) HandleUpdateShortURL(params urls.UpdateURLParams, principal any) middleware.Responder {
	//Validate params
	err := validateParams(params)
	if err != nil {
		return urls.NewUpdateURLBadRequest().WithPayload(getError(err))
	}

	//Check if the account API key is valid
	// TODO
	account, err := handlers.AccountRepository.GetAccountByApiKey(context.Background(), principal.(string))
	if err != nil {
		if err == sql.ErrNoRows {
			return urls.NewUpdateURLUnauthorized().WithPayload(&models.Error{
				Error: "Invalid API Key"})
		}
		return internalErrorInUpdateShortURL(err)
	}

	//Check if the customer is active
	if !account.Enabled {
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
			return urls.NewUpdateURLNotFound().WithPayload(&models.Error{
				Error: "Short URL Not Found"})
		}
		return internalErrorInUpdateShortURL(err)
	}
	if params.Body.NewLongURL != "" {
		url.LongUrl = params.Body.NewLongURL
		err = handlers.URLRepository.UpdateLongURL(context.Background(), db.UpdateLongURLParams{
			LongUrl: params.Body.NewLongURL,
			ID:      url.ID,
		})
		if err != nil {
			return internalErrorInUpdateShortURL(err)
		}
	}
	if params.Body.Status != "" {
		url.Enabled = encodeStatus(params.Body.Status)
		err = handlers.URLRepository.UpdateURLStatus(context.Background(), db.UpdateURLStatusParams{
			Enabled: url.Enabled,
			ID:      url.ID,
		})
		if err != nil {
			return internalErrorInUpdateShortURL(err)
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
			return internalErrorInUpdateShortURL(err)
		}
	}
	return urls.NewUpdateURLOK().WithPayload(&urls.UpdateURLOKBody{
		LongURL:        url.LongUrl,
		Status:         decodeStatus(url.Enabled),
		TrackingStatus: decodeStatus(url.TrackingEnabled),
	})
}

func internalErrorInUpdateShortURL(err error) middleware.Responder {
	log.Println("Error updating short URL: ", err)
	return urls.NewUpdateURLInternalServerError().WithPayload(&models.Error{
		Error: err.Error()})
}

func validateParams(params urls.UpdateURLParams) error {
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
	// TODO
	// if params.XAPIKEY == "" {
	// 	return errors.New("missing x-api-key header")
	// }
	return nil
}
