package core

import (
	"context"
	"database/sql"
	"log/slog"
	"net/url"

	"github.com/elbombardi/squrl/src/db"
	"github.com/elbombardi/squrl/src/redirection_service/util"
)

type LinksService struct {
	db.LinkRepository
	db.AccountRepository
	db.ClickRepository
	*util.Config
	*slog.Logger
}

func (s *LinksService) Resolve(params *ResolveLinkParams) (*Link, error) {

	if params == nil {
		return nil, CoreError{Code: ErrBadParams, Message: "Bad parameters"}
	}

	// Retrieve account information from the database
	account, err := s.getAccountInfo(params.AccountPrefix)
	if err != nil {
		if err == sql.ErrNoRows {
			s.Error("Account not found", "Account", params.AccountPrefix)
			return nil, CoreError{Code: ErrAccountNotFound, Message: "Account not found"}
		}
		s.Error("Unexpected error while retrieving account information", "Details", err)
		return nil, err
	}

	// Check if the account is enabled
	if !account.Enabled {
		s.Info("Click on a link belonging to a disabled account", "URL", params.ShortUrl, "Account", account.Username)
		return nil, CoreError{Code: ErrAccountDisabled, Message: "Account disabled"}
	}

	// Retrieve Short link information from the database
	link, err := s.getLinkInfo(account.ID, params.ShortUrlKey)
	if err != nil {
		if err == sql.ErrNoRows {
			s.Error("Link does not exist", "URL", params.ShortUrl)
			return nil, CoreError{Code: ErrLinkNotFound, Message: "Short URL not found"}
		}
		s.Error("Unexpected error retrieving short URL information", "Details", err)
		return nil, err
	}

	// Check if the link is enabled
	if !link.Enabled {
		s.Info("Click on a disabled URL", "URL", params.ShortUrl)
		return nil, CoreError{Code: ErrLinkDisabled, Message: "Short URL disabled"}
	}
	link.Account = account

	// Persist click information
	if link.TrackingEnabled {
		err = s.InsertNewClick(context.Background(), db.InsertNewClickParams{
			LinkID:    link.ID,
			UserAgent: sql.NullString{String: params.UserAgent, Valid: true},
			IpAddress: sql.NullString{String: params.IpAddress, Valid: true},
		})
		if err != nil {
			s.Error("Unexptected error while registering a new click", "Details", err)
			return nil, err
		}
	}

	// Return resolved link
	s.Info("URL Resolved", "URL", params.ShortUrl)
	return link, nil
}

func (s *LinksService) getAccountInfo(prefix string) (*Account, error) {
	dbAcc, err := s.GetAccountByPrefix(context.Background(), prefix)
	if err != nil {
		return nil, err
	}
	return &Account{
		ID:       dbAcc.ID,
		Prefix:   dbAcc.Prefix,
		Username: dbAcc.Username,
		Enabled:  dbAcc.Enabled,
	}, nil
}

func (s *LinksService) getLinkInfo(accountId int32, key string) (*Link, error) {
	dbLink, err := s.GetLinkByAccountIDAndShortURLKey(context.Background(),
		db.GetLinkByAccountIDAndShortURLKeyParams{
			AccountID: accountId,
			ShortUrlKey: sql.NullString{
				String: key,
				Valid:  true,
			},
		},
	)
	if err != nil {
		return nil, err
	}
	parsedUrl, _ := url.Parse(dbLink.LongUrl)
	return &Link{
		ID:              dbLink.ID,
		LongUrl:         *parsedUrl,
		Enabled:         dbLink.Enabled,
		TrackingEnabled: dbLink.TrackingEnabled,
	}, nil
}
