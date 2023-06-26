package routes

import (
	db "github.com/elbombardi/squrl/db/sqlc"
)

type Routes struct {
	CustomersRepository db.CustomersRepository
	ShortURLsRepository db.ShortURLsRepository
	ClicksRepository    db.ClicksRepository
}
