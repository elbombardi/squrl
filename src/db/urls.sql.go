// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: urls.sql

package db

import (
	"context"
	"database/sql"
)

const checkShortUrlKeyExists = `-- name: CheckShortUrlKeyExists :one
SELECT EXISTS(SELECT 1 FROM url WHERE short_url_key = $1 AND account_id = $2)
`

type CheckShortUrlKeyExistsParams struct {
	ShortUrlKey sql.NullString
	AccountID   int32
}

func (q *Queries) CheckShortUrlKeyExists(ctx context.Context, arg CheckShortUrlKeyExistsParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkShortUrlKeyExists, arg.ShortUrlKey, arg.AccountID)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getURLByAccountIDAndShortURLKey = `-- name: GetURLByAccountIDAndShortURLKey :one
SELECT id, short_url_key, account_id, long_url, enabled, tracking_enabled,created_at, updated_at
FROM url WHERE account_id = $1 AND short_url_key = $2
`

type GetURLByAccountIDAndShortURLKeyParams struct {
	AccountID   int32
	ShortUrlKey sql.NullString
}

func (q *Queries) GetURLByAccountIDAndShortURLKey(ctx context.Context, arg GetURLByAccountIDAndShortURLKeyParams) (Url, error) {
	row := q.db.QueryRowContext(ctx, getURLByAccountIDAndShortURLKey, arg.AccountID, arg.ShortUrlKey)
	var i Url
	err := row.Scan(
		&i.ID,
		&i.ShortUrlKey,
		&i.AccountID,
		&i.LongUrl,
		&i.Enabled,
		&i.TrackingEnabled,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertNewURL = `-- name: InsertNewURL :exec
INSERT INTO url (short_url_key, account_id, long_url)
VALUES ($1, $2, $3)
RETURNING id, short_url_key, account_id, long_url, enabled, tracking_enabled, created_at, updated_at
`

type InsertNewURLParams struct {
	ShortUrlKey sql.NullString
	AccountID   int32
	LongUrl     string
}

func (q *Queries) InsertNewURL(ctx context.Context, arg InsertNewURLParams) error {
	_, err := q.db.ExecContext(ctx, insertNewURL, arg.ShortUrlKey, arg.AccountID, arg.LongUrl)
	return err
}

const updateLongURL = `-- name: UpdateLongURL :exec
UPDATE url SET long_url = $1, updated_at=now() WHERE id = $2
RETURNING id, short_url_key, account_id, long_url, enabled, tracking_enabled, created_at, updated_at
`

type UpdateLongURLParams struct {
	LongUrl string
	ID      int32
}

func (q *Queries) UpdateLongURL(ctx context.Context, arg UpdateLongURLParams) error {
	_, err := q.db.ExecContext(ctx, updateLongURL, arg.LongUrl, arg.ID)
	return err
}

const updateURLStatus = `-- name: UpdateURLStatus :exec
UPDATE url SET enabled = $1, updated_at=now() WHERE id = $2 
RETURNING id, short_url_key, account_id, long_url, enabled, tracking_enabled, created_at, updated_at
`

type UpdateURLStatusParams struct {
	Enabled bool
	ID      int32
}

func (q *Queries) UpdateURLStatus(ctx context.Context, arg UpdateURLStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateURLStatus, arg.Enabled, arg.ID)
	return err
}

const updateURLTrackingStatus = `-- name: UpdateURLTrackingStatus :exec
UPDATE url SET tracking_enabled = $1, updated_at=now() WHERE id = $2 
RETURNING id, short_url_key, account_id, long_url, enabled, tracking_enabled, created_at, updated_at
`

type UpdateURLTrackingStatusParams struct {
	TrackingEnabled bool
	ID              int32
}

func (q *Queries) UpdateURLTrackingStatus(ctx context.Context, arg UpdateURLTrackingStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateURLTrackingStatus, arg.TrackingEnabled, arg.ID)
	return err
}
