package handlers

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/elbombardi/squrl/src/api_service/api/operations"
	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
	"github.com/go-openapi/runtime/middleware"
)

func (h *Handlers) HandleUpdateShortURL(params operations.PutShortURLParams) middleware.Responder {
	//Validate params
	err := validateUpdateShortURLParams(params)
	if err != nil {
		return operations.NewPutShortURLBadRequest().WithPayload(&operations.PutShortURLBadRequestBody{
			Error: err.Error()})
	}
	//Check if the account API key is valid
	account, err := h.AccountRepository.GetAccountByApiKey(context.Background(), params.XAPIKEY)
	if err != nil {
		if err == sql.ErrNoRows {
			return operations.NewPostShortURLUnauthorized().WithPayload(&operations.PostShortURLUnauthorizedBody{
				Error: "Invalid API Key"})
		}
		return internalErrorInUpdateShortURL(err)
	}
	//Check if the customer is active
	if !account.Enabled {
		return operations.NewPostShortURLUnauthorized().WithPayload(&operations.PostShortURLUnauthorizedBody{
			Error: "Account disabled"})
	}

	url, err := h.URLRepository.GetURLByAccountIDAndShortURLKey(context.Background(),
		db.GetURLByAccountIDAndShortURLKeyParams{
			AccountID: account.ID,
			ShortUrlKey: sql.NullString{
				String: *params.Body.ShortURLKey,
				Valid:  true,
			},
		})
	if err != nil {
		if err == sql.ErrNoRows {
			return operations.NewPutShortURLNotFound().WithPayload(&operations.PutShortURLNotFoundBody{
				Error: "Short URL Not Found"})
		}
		return internalErrorInUpdateShortURL(err)
	}
	if params.Body.NewLongURL != "" {
		url.LongUrl = params.Body.NewLongURL
		err = h.URLRepository.UpdateLongURL(context.Background(), db.UpdateLongURLParams{
			LongUrl: params.Body.NewLongURL,
			ID:      url.ID,
		})
		if err != nil {
			return internalErrorInUpdateShortURL(err)
		}
	}
	if params.Body.Status != "" {
		url.Enabled = encodeStatus(params.Body.Status)
		err = h.URLRepository.UpdateURLStatus(context.Background(), db.UpdateURLStatusParams{
			Enabled: url.Enabled,
			ID:      url.ID,
		})
		if err != nil {
			return internalErrorInUpdateShortURL(err)
		}
	}
	if params.Body.TrackingStatus != "" {
		url.TrackingEnabled = encodeStatus(params.Body.TrackingStatus)
		err = h.URLRepository.UpdateURLTrackingStatus(context.Background(),
			db.UpdateURLTrackingStatusParams{
				TrackingEnabled: url.TrackingEnabled,
				ID:              url.ID,
			})
		if err != nil {
			return internalErrorInUpdateShortURL(err)
		}
	}
	return operations.NewPutShortURLOK().WithPayload(&operations.PutShortURLOKBody{
		LongURL:        url.LongUrl,
		Status:         decodeStatus(url.Enabled),
		TrackingStatus: decodeStatus(url.TrackingEnabled),
	})
}

func internalErrorInUpdateShortURL(err error) middleware.Responder {
	log.Println("Error updating short URL: ", err)
	return operations.NewPutShortURLInternalServerError().WithPayload(&operations.PutShortURLInternalServerErrorBody{
		Error: err.Error()})
}

func validateUpdateShortURLParams(params operations.PutShortURLParams) error {
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
	if params.XAPIKEY == "" {
		return errors.New("missing x-api-key header")
	}
	return nil
}
