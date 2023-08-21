package core

import "net/url"

const (
	ErrUnauthorized = iota
	ErrBadParams
	ErrAccountNotFound
	ErrAccountDisabled
	ErrLinkNotFound
	ErrLinkDisabled
)

type CoreError struct {
	Code    int
	Message string
}

func (e CoreError) Error() string { return e.Message }

// Interface that provides link management methods.
type LinksManager interface {
	// Resolves a link by its short url key and the account's prefix.
	//
	// Returns the link if it exists, otherwise it returns an error.
	Resolve(params *ResolveLinkParams) (*Link, error)
}

type ResolveLinkParams struct {
	AccountPrefix string
	ShortUrlKey   string
	UserAgent     string
	IpAddress     string
	ShortUrl      string
}

type Account struct {
	ID       int32
	Prefix   string
	Username string
	Enabled  bool
}

type Link struct {
	*Account
	ID              int32
	LongUrl         url.URL
	Enabled         bool
	TrackingEnabled bool
}
