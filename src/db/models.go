// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"database/sql"
)

// Table holding Account information
type Account struct {
	ID int32
	// 3 characters, case-sensitive
	Prefix   string
	Username string
	Email    string
	// Hashed password
	HashedPassword string
	// A flag that enables/disables the account and its urls
	Enabled bool
	// Timestamp of creation
	CreatedAt sql.NullTime
	// Timestamp of last update
	UpdatedAt sql.NullTime
}

// Table holding click information
type Click struct {
	ID     int32
	LinkID int32
	// Timestamp of click
	ClickDateTime sql.NullTime
	UserAgent     sql.NullString
	IpAddress     sql.NullString
}

// Table holding Link information
type Link struct {
	ID int32
	// 6 characters, case-sensitive
	ShortUrlKey sql.NullString
	AccountID   int32
	LongUrl     string
	// A flag to enable/disable the url redirection
	Enabled bool
	// A flag that enable/disable url tracking on redirection
	TrackingEnabled bool
	// Timestamp of creation
	CreatedAt sql.NullTime
	// Timestamp of last update
	UpdatedAt sql.NullTime
}
