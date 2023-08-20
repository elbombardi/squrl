package core

import "net/url"

const (
	ERR_UNAUTHORIZED = iota
	ERR_BAD_PARAMS
	ERR_ACCOUNT_NOT_FOUND
	ERR_ACCOUNT_DISABLED
	ERR_LINK_NOT_FOUND
	ERR_LINK_DISABLED
)

type CoreError struct {
	Code    int
	Message string
}

func (e CoreError) Error() string { return e.Message }

// Interface that provides authentication methods.
type Authenticator interface {
	// Authenticate takes a username and password
	//
	// Returns a token if the credentials are valid.
	Authenticate(username string, password string) (token string, err error)

	// Validate takes a token
	//
	// Returns the User if the token is valid Otherwise, it returns an error.
	Validate(token string) (user *User, err error)
}

// Interface that provides account management methods.
type AccountsManager interface {
	// CreateAccount creates a new account.
	//
	// The user parameter must be an admin
	// The response contains the password and the prefix for the new account
	Create(params *CreateAccountParams, user *User) (*CreateAccountResponse, error)

	// UpdateAccount updates an existing account (enable/disable an account).
	//
	// The user parameter must be an admin
	// The response contains the new enabled status
	Update(params *UpdateAccountParams, user *User) (*UpdateAccountResponse, error)
}

// Interface that provides link management methods.
type LinksManager interface {
	// Creates a new short url.
	//
	// The user parameter must be a valid account username
	// The response contains the short url key and the short url
	Shorten(longUrl string, user *User) (*Link, error)

	// Updates an existing link.
	//
	// The user parameter must be a valid account username
	// The response contains the up to date version of the short url information
	Update(params *LinkUpdateParams, user *User) (*Link, error)
}

type User struct {
	Username string
	IsAdmin  bool
}

func (u *User) String() string {
	return u.Username
}

type CreateAccountParams struct {
	Email    string
	Username string
}

type CreateAccountResponse struct {
	Password string
	Prefix   string
}

type UpdateAccountParams struct {
	Enabled  bool
	Username string
}

type UpdateAccountResponse struct {
	Enabled bool
}

type Link struct {
	LongUrl         url.URL
	ShortUrlKey     string
	ShortUrl        url.URL
	Enabled         bool
	TrackingEnabled bool
}

type LinkUpdateParams struct {
	ShortUrlKey     string
	NewLongURL      Optional[string]
	Enabled         Optional[bool]
	TrackingEnabled Optional[bool]
}

type Optional[T any] struct {
	Value T
	IsSet bool
}
