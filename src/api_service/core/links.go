package core

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/url"

	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
)

type LinksService struct {
	db.URLRepository
	db.AccountRepository
	*util.Config
}

func (service *LinksService) Shorten(longUrl string, user *User) (*Link, error) {

	// Check if the user is authenticated
	if user == nil {
		return nil, CoreError{
			Code:    ERR_UNAUTHORIZED,
			Message: "Unauthorized access",
		}
	}

	// Validate params
	parsedUrl, err := validateUrl(longUrl)
	if err != nil {
		slog.Error("Shorten URL bad params", "Details", err)
		return nil, CoreError{
			Code:    ERR_BAD_PARAMS,
			Message: err.Error(),
		}
	}

	// Check if the account exists
	account, err := service.AccountRepository.GetAccountByUsername(context.Background(), user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("Account not found", "Account", user)
			return nil, CoreError{
				Code:    ERR_ACCOUNT_NOT_FOUND,
				Message: "Account not found",
			}
		}
		slog.Error("Unexpected error while retrieving account by username", "Username", user, "Detail", err)
		return nil, err
	}

	// Check if the account is active
	if !account.Enabled {
		slog.Info("Account disabled", "Account", user)
		return nil, CoreError{
			Code:    ERR_ACCOUNT_DISABLED,
			Message: "Account disabled",
		}
	}

	// Generate short URL key
	shortURLKey, err := service.generateShortURLKey(&account)
	if err != nil {
		slog.Error("Unexpected error while generating short URL", "Details", err)
		return nil, err
	}

	// Insert the new short URL in the database
	err = service.URLRepository.InsertNewURL(context.Background(), db.InsertNewURLParams{
		ShortUrlKey: sql.NullString{String: shortURLKey, Valid: true},
		LongUrl:     longUrl,
		AccountID:   account.ID,
	})
	if err != nil {
		slog.Error("Error inserting new URL in DB", "Details", err)
		return nil, err
	}

	slog.Info("New short URL created successfully", "Account", user, "LongUrl", longUrl)
	shortUrl := service.buildShortUrl(&account, shortURLKey)

	// Return the short URL
	return &Link{
		LongUrl:         *parsedUrl,
		ShortUrl:        shortUrl,
		ShortUrlKey:     shortURLKey,
		Enabled:         true,
		TrackingEnabled: true,
	}, nil
}

func (service *LinksService) generateShortURLKey(account *db.Account) (string, error) {
	shortURLKey := util.GenerateRandomString(6)

	// Check if the url key is unique
	exists, err := service.URLRepository.CheckShortUrlKeyExists(context.Background(), db.CheckShortUrlKeyExistsParams{
		ShortUrlKey: sql.NullString{String: shortURLKey, Valid: true},
		AccountID:   account.ID,
	})
	if err != nil {
		return "", err
	}

	if exists {
		return service.generateShortURLKey(account)
	}

	return shortURLKey, nil
}

func (service *LinksService) Update(params *LinkUpdateParams, user *User) (*Link, error) {

	// Check if the user is authenticated
	if user == nil {
		return nil, CoreError{
			Code:    ERR_UNAUTHORIZED,
			Message: "unauthorized access",
		}
	}

	//Validate params
	parsedUrl, err := validateParams(params)
	if err != nil {
		slog.Error("Bad UpdateURL params", "Details", err)
		return nil, CoreError{
			Code:    ERR_BAD_PARAMS,
			Message: err.Error(),
		}
	}

	// Check if the account exists
	account, err := service.AccountRepository.GetAccountByUsername(context.Background(), user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("Account not found", "Account", user)
			return nil, CoreError{
				Code:    ERR_ACCOUNT_NOT_FOUND,
				Message: "account not found",
			}
		}
		slog.Error("Unexpected error while retrieving account by username",
			"Username", user, "Details", err)
		return nil, err
	}

	// Check if the account is active
	if !account.Enabled {
		slog.Info("Account disabled", "Account", account.Username)
		return nil, CoreError{
			Code:    ERR_ACCOUNT_DISABLED,
			Message: "account disabled",
		}
	}

	savedLink, err := service.URLRepository.GetURLByAccountIDAndShortURLKey(context.Background(),
		db.GetURLByAccountIDAndShortURLKeyParams{
			AccountID: account.ID,
			ShortUrlKey: sql.NullString{
				String: params.ShortUrlKey,
				Valid:  true,
			},
		})
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("URL not found", "Account", account.Username, "ShortURLKey", params.ShortUrlKey)
			return nil, CoreError{
				Code:    ERR_LINK_NOT_FOUND,
				Message: "url not found",
			}
		}
		slog.Error("Unxpected error while retrieving short URL", "Account", account.Username,
			"ShortURLKey", params.ShortUrlKey, "Details", err)
		return nil, err
	}

	if params.NewLongURL.IsSet {
		savedLink.LongUrl = params.NewLongURL.Value
		err = service.URLRepository.UpdateLongURL(context.Background(), db.UpdateLongURLParams{
			LongUrl: params.NewLongURL.Value,
			ID:      savedLink.ID,
		})
		if err != nil {
			slog.Error("Unexpected error while updating long URL", "Account", account.Username,
				"ShortURLKey", params.ShortUrlKey, "Details", err)
			return nil, err
		}
	}

	if params.Enabled.IsSet {
		savedLink.Enabled = params.Enabled.Value
		err = service.URLRepository.UpdateURLStatus(context.Background(), db.UpdateURLStatusParams{
			Enabled: params.Enabled.Value,
			ID:      savedLink.ID,
		})
		if err != nil {
			slog.Error("Unexpected error while updating URL enabled status", "Account", account.Username,
				"ShortURLKey", params.ShortUrlKey, "Details", err)
			return nil, err
		}
	}

	if params.TrackingEnabled.IsSet {
		savedLink.TrackingEnabled = params.TrackingEnabled.Value
		err = service.URLRepository.UpdateURLTrackingStatus(context.Background(),
			db.UpdateURLTrackingStatusParams{
				TrackingEnabled: params.TrackingEnabled.Value,
				ID:              savedLink.ID,
			})
		if err != nil {
			slog.Error("Unexpected error while updating URL's tracking enabled status", "Account", account.Username,
				"ShortURLKey", params.ShortUrlKey, "Details", err)
			return nil, err
		}
	}

	if parsedUrl == nil {
		parsedUrl, _ = url.Parse(savedLink.LongUrl)
	}

	slog.Info("URL updated successfully", "Account", account.Username, "Params", *params)
	return &Link{
		ShortUrlKey:     params.ShortUrlKey,
		LongUrl:         *parsedUrl,
		Enabled:         savedLink.Enabled,
		TrackingEnabled: savedLink.TrackingEnabled,
	}, nil
}

func validateParams(params *LinkUpdateParams) (*url.URL, error) {
	if params == nil {
		return nil, errors.New("missing parameter : 'params'")
	}
	if params.ShortUrlKey == "" {
		return nil, errors.New("missing parameter : 'short_url_key'")
	}
	var parsedUrl *url.URL
	var err error
	if params.NewLongURL.IsSet {
		parsedUrl, err = validateUrl(params.NewLongURL.Value)
		if err != nil {
			return nil, err
		}
	}
	return parsedUrl, nil
}

func validateUrl(rawUrl string) (*url.URL, error) {

	// URL is required
	if rawUrl == "" {
		return nil, errors.New("Url is empty")
	}

	parsedUrl, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}

	if parsedUrl.Scheme != "http" && parsedUrl.Scheme != "https" {
		return nil, errors.New("Url scheme is not http or https")
	}

	return parsedUrl, nil
}

func (service *LinksService) buildShortUrl(account *db.Account, shortUrlKey string) url.URL {
	shortUrl, _ := url.Parse(fmt.Sprintf("%v/%v/%v", service.Config.RedirectionBaseURL, account.Prefix, shortUrlKey))
	return *shortUrl
}
