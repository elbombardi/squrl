// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: clicks.sql

package db

import (
	"context"
	"database/sql"
)

const insertNewClick = `-- name: InsertNewClick :exec
INSERT INTO click (link_id, user_agent, ip_address)
VALUES ($1, $2, $3)
`

type InsertNewClickParams struct {
	LinkID    int32
	UserAgent sql.NullString
	IpAddress sql.NullString
}

func (q *Queries) InsertNewClick(ctx context.Context, arg InsertNewClickParams) error {
	_, err := q.db.ExecContext(ctx, insertNewClick, arg.LinkID, arg.UserAgent, arg.IpAddress)
	return err
}
