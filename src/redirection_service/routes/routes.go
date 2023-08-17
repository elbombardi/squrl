package routes

import "github.com/elbombardi/squrl/src/db"

type Routes struct {
	db.AccountRepository
	db.URLRepository
	db.ClickRepository
}
