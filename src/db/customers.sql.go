// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: customers.sql

package db

import (
	"context"
)

const checkApiKeyExists = `-- name: CheckApiKeyExists :one
SELECT EXISTS(SELECT 1 FROM account WHERE api_key = $1)
`

func (q *Queries) CheckApiKeyExists(ctx context.Context, apiKey string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkApiKeyExists, apiKey)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const checkPrefixExists = `-- name: CheckPrefixExists :one
SELECT EXISTS(SELECT 1 FROM account WHERE prefix = $1)
`

func (q *Queries) CheckPrefixExists(ctx context.Context, prefix string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkPrefixExists, prefix)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const checkUsernameExists = `-- name: CheckUsernameExists :one
SELECT EXISTS(SELECT 1 FROM account WHERE username = $1)
`

func (q *Queries) CheckUsernameExists(ctx context.Context, username string) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkUsernameExists, username)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getAccountByApiKey = `-- name: GetAccountByApiKey :one
SELECT id, prefix, username, email, api_key, enabled, created_at, updated_at
FROM account WHERE api_key = $1
`

func (q *Queries) GetAccountByApiKey(ctx context.Context, apiKey string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByApiKey, apiKey)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Prefix,
		&i.Username,
		&i.Email,
		&i.ApiKey,
		&i.Enabled,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAccountByPrefix = `-- name: GetAccountByPrefix :one
SELECT id, prefix, username, email, api_key, enabled, created_at, updated_at
FROM account WHERE prefix = $1
`

func (q *Queries) GetAccountByPrefix(ctx context.Context, prefix string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByPrefix, prefix)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Prefix,
		&i.Username,
		&i.Email,
		&i.ApiKey,
		&i.Enabled,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAccountByUsername = `-- name: GetAccountByUsername :one
SELECT id, prefix, username, email, api_key, enabled, created_at, updated_at
FROM account WHERE username = $1
`

func (q *Queries) GetAccountByUsername(ctx context.Context, username string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccountByUsername, username)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Prefix,
		&i.Username,
		&i.Email,
		&i.ApiKey,
		&i.Enabled,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertNewAccount = `-- name: InsertNewAccount :exec
INSERT INTO account (prefix, username, email, api_key)
VALUES ($1, $2, $3, $4)
RETURNING id, prefix, username, email, api_key, enabled, created_at, updated_at
`

type InsertNewAccountParams struct {
	Prefix   string
	Username string
	Email    string
	ApiKey   string
}

func (q *Queries) InsertNewAccount(ctx context.Context, arg InsertNewAccountParams) error {
	_, err := q.db.ExecContext(ctx, insertNewAccount,
		arg.Prefix,
		arg.Username,
		arg.Email,
		arg.ApiKey,
	)
	return err
}

const updateAccountStatusByUsername = `-- name: UpdateAccountStatusByUsername :exec
UPDATE account SET enabled = $1, updated_at=now() WHERE username = $2
RETURNING id, prefix, username, email, api_key, enabled, created_at, updated_at
`

type UpdateAccountStatusByUsernameParams struct {
	Enabled  bool
	Username string
}

func (q *Queries) UpdateAccountStatusByUsername(ctx context.Context, arg UpdateAccountStatusByUsernameParams) error {
	_, err := q.db.ExecContext(ctx, updateAccountStatusByUsername, arg.Enabled, arg.Username)
	return err
}