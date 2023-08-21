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
	db.LinkRepository
	db.AccountRepository
	*util.Config
	*slog.Logger
}

func (s *LinksService) Shorten(longUrl string, user *User) (*Link, error) {

	// Check if the user is authenticated
	if user == nil {
		return nil, CoreError{
			Code:    ErrUnauthorized,
			Message: "Unauthorized access",
		}
	}

	// Validate params
	parsedUrl, err := validateUrl(longUrl)
	if err != nil {
		s.Error("Shorten URL bad params", "Details", err)
		return nil, CoreError{
			Code:    ErrBadParams,
			Message: err.Error(),
		}
	}

	// Check if the account exists
	account, err := s.GetAccountByUsername(context.Background(), user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			s.Error("Account not found", "Account", user)
			return nil, CoreError{
				Code:    ErrAccountNotFound,
				Message: "Account not found",
			}
		}
		s.Error("Unexpected error while retrieving account by username", "Username", user, "Detail", err)
		return nil, err
	}

	// Check if the account is active
	if !account.Enabled {
		s.Error("Account is disabled and cannot be used to access the service", "Account", user)
		return nil, CoreError{
			Code:    ErrAccountDisabled,
			Message: "Account disabled",
		}
	}

	// Generate short URL key
	tryAgain := true
	var shortURLKey string
	for tryAgain {
		tryAgain, shortURLKey, err = s.generateShortURLKey(&account)
		if err != nil {
			s.Error("Unexpected error while generating short URL", "Details", err)
			return nil, err
		}
	}

	// Insert the new short URL in the database
	err = s.InsertNewLink(context.Background(), db.InsertNewLinkParams{
		ShortUrlKey: sql.NullString{String: shortURLKey, Valid: true},
		LongUrl:     longUrl,
		AccountID:   account.ID,
	})
	if err != nil {
		s.Error("Error inserting new URL in DB", "Details", err)
		return nil, err
	}

	s.Info("New short URL created successfully", "Account", user, "LongUrl", longUrl)
	shortUrl := s.buildShortUrl(&account, shortURLKey)
	// Return the short URL
	return &Link{
		LongUrl:         *parsedUrl,
		ShortUrl:        shortUrl,
		ShortUrlKey:     shortURLKey,
		Enabled:         true,
		TrackingEnabled: true,
	}, nil
}

func (s *LinksService) generateShortURLKey(account *db.Account) (bool, string, error) {
	shortURLKey := util.GenerateRandomString(6)

	// Check if the url key is unique
	exists, err := s.CheckShortUrlKeyExists(context.Background(), db.CheckShortUrlKeyExistsParams{
		ShortUrlKey: sql.NullString{String: shortURLKey, Valid: true},
		AccountID:   account.ID,
	})
	if err != nil {
		return false, "", err
	}

	return exists, shortURLKey, nil
}

func (s *LinksService) Update(params *LinkUpdateParams, user *User) (*Link, error) {

	// Check if the user is authenticated
	if user == nil {
		return nil, CoreError{
			Code:    ErrUnauthorized,
			Message: "unauthorized access",
		}
	}

	//Validate params
	parsedUrl, err := validateParams(params)
	if err != nil {
		s.Error("Bad UpdateURL params", "Details", err)
		return nil, CoreError{
			Code:    ErrBadParams,
			Message: err.Error(),
		}
	}

	// Check if the account exists
	account, err := s.GetAccountByUsername(context.Background(), user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			s.Error("Account not found", "Account", user)
			return nil, CoreError{
				Code:    ErrAccountNotFound,
				Message: "account not found",
			}
		}
		s.Error("Unexpected error while retrieving account by username",
			"Username", user, "Details", err)
		return nil, err
	}

	// Check if the account is active
	if !account.Enabled {
		s.Error("Account is disabled, and cannot be used to access the service", "Account", account.Username)
		return nil, CoreError{
			Code:    ErrAccountDisabled,
			Message: "account disabled",
		}
	}

	savedLink, err := s.GetLinkByAccountIDAndShortURLKey(context.Background(),
		db.GetLinkByAccountIDAndShortURLKeyParams{
			AccountID: account.ID,
			ShortUrlKey: sql.NullString{
				String: params.ShortUrlKey,
				Valid:  true,
			},
		})
	if err != nil {
		if err == sql.ErrNoRows {
			s.Error("URL not found", "Account", account.Username, "ShortURLKey", params.ShortUrlKey)
			return nil, CoreError{
				Code:    ErrLinkNotFound,
				Message: "url not found",
			}
		}
		s.Error("Unxpected error while retrieving short URL", "Account", account.Username,
			"ShortURLKey", params.ShortUrlKey, "Details", err)
		return nil, err
	}

	if params.NewLongURL.IsSet {
		savedLink.LongUrl = params.NewLongURL.Value
		err = s.UpdateLinkLongURL(context.Background(), db.UpdateLinkLongURLParams{
			LongUrl: params.NewLongURL.Value,
			ID:      savedLink.ID,
		})
		if err != nil {
			s.Error("Unexpected error while updating long URL", "Account", account.Username,
				"ShortURLKey", params.ShortUrlKey, "Details", err)
			return nil, err
		}
	}

	if params.Enabled.IsSet {
		savedLink.Enabled = params.Enabled.Value
		err = s.UpdateLinkStatus(context.Background(), db.UpdateLinkStatusParams{
			Enabled: params.Enabled.Value,
			ID:      savedLink.ID,
		})
		if err != nil {
			s.Error("Unexpected error while updating URL enabled status", "Account", account.Username,
				"ShortURLKey", params.ShortUrlKey, "Details", err)
			return nil, err
		}
	}

	if params.TrackingEnabled.IsSet {
		savedLink.TrackingEnabled = params.TrackingEnabled.Value
		err = s.UpdateLinkTrackingStatus(context.Background(),
			db.UpdateLinkTrackingStatusParams{
				TrackingEnabled: params.TrackingEnabled.Value,
				ID:              savedLink.ID,
			})
		if err != nil {
			s.Error("Unexpected error while updating URL's tracking enabled status", "Account", account.Username,
				"ShortURLKey", params.ShortUrlKey, "Details", err)
			return nil, err
		}
	}

	if parsedUrl == nil {
		parsedUrl, _ = url.Parse(savedLink.LongUrl)
	}

	s.Info("URL updated successfully", "Account", account.Username, "Params", *params)
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

func (s *LinksService) buildShortUrl(account *db.Account, shortUrlKey string) url.URL {
	shortUrl, _ := url.Parse(fmt.Sprintf("%v/%v/%v", s.RedirectionBaseURL, account.Prefix, shortUrlKey))
	return *shortUrl
}
