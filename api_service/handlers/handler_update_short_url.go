package handlers

import (
	"context"
	"database/sql"
	"errors"

	"github.com/elbombardi/squrl/api_service/api/operations"
	db "github.com/elbombardi/squrl/db/sqlc"
	"github.com/elbombardi/squrl/util"
	"github.com/go-openapi/runtime/middleware"
)

func (h *Handlers) UpdateShortURLHandler(params operations.PutShortURLParams) middleware.Responder {
	//Validate params
	err := validateUpdateShortURLParams(params)
	if err != nil {
		return operations.NewPutShortURLBadRequest().WithPayload(&operations.PutShortURLBadRequestBody{
			Error: err.Error()})
	}
	//Check if the customer API key is valid
	customer, err := h.CustomersRepository.GetCustomerByApiKey(context.Background(), params.XAPIKEY)
	if err != nil {
		if err == sql.ErrNoRows {
			return operations.NewPostShortURLUnauthorized().WithPayload(&operations.PostShortURLUnauthorizedBody{
				Error: "Invalid API Key"})
		}
		return internalErrorInUpdateShortURL(err)
	}
	//Check if the customer is active
	if customer.Status != "e" {
		return operations.NewPostShortURLUnauthorized().WithPayload(&operations.PostShortURLUnauthorizedBody{
			Error: "Customer disabled"})
	}

	shortUrl, err := h.ShortURLsRepository.GetShortURLByCustomerIDAndShortURLKey(context.Background(),
		db.GetShortURLByCustomerIDAndShortURLKeyParams{
			CustomerID: customer.ID,
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
		shortUrl.LongUrl = params.Body.NewLongURL
		err = h.ShortURLsRepository.UpdateShortURLLongURL(context.Background(), db.UpdateShortURLLongURLParams{
			LongUrl: params.Body.NewLongURL,
			ID:      shortUrl.ID,
		})
		if err != nil {
			return internalErrorInUpdateShortURL(err)
		}
	}
	if params.Body.Status != "" {
		newStatus := sql.NullString{
			String: encodeStatus(params.Body.Status),
			Valid:  true,
		}
		shortUrl.Status = newStatus
		err = h.ShortURLsRepository.UpdateShortURLStatus(context.Background(), db.UpdateShortURLStatusParams{
			Status: newStatus,
			ID:     shortUrl.ID,
		})
		if err != nil {
			return internalErrorInUpdateShortURL(err)
		}
	}
	if params.Body.TrackingStatus != "" {
		newTrackingStatus := sql.NullString{
			String: encodeStatus(params.Body.TrackingStatus),
			Valid:  true,
		}
		shortUrl.Status = newTrackingStatus
		err = h.ShortURLsRepository.UpdateShortURLTrackingStatus(context.Background(),
			db.UpdateShortURLTrackingStatusParams{
				TrackingStatus: newTrackingStatus,
				ID:             shortUrl.ID,
			})
		if err != nil {
			return internalErrorInUpdateShortURL(err)
		}
	}
	return operations.NewPutShortURLOK().WithPayload(&operations.PutShortURLOKBody{
		LongURL:        shortUrl.LongUrl,
		Status:         decodeStatus(shortUrl.Status.String),
		TrackingStatus: decodeStatus(shortUrl.TrackingStatus.String),
	})
}

func internalErrorInUpdateShortURL(err error) middleware.Responder {
	return operations.NewPutShortURLInternalServerError().WithPayload(&operations.PutShortURLInternalServerErrorBody{
		Error: err.Error()})
}

func validateUpdateShortURLParams(params operations.PutShortURLParams) error {
	if params.Body.ShortURLKey == nil {
		return errors.New("missing parameter : 'short_url_key'")
	}
	if params.Body.NewLongURL != "" {
		if valid, err := util.IsValidURL(params.Body.NewLongURL); !valid {
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
