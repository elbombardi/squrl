package main

import (
	"log"

	db "github.com/elbombardi/squrl/db/sqlc"
)

func finalizeApp() {
	defer finalizeDatabase()
}

func finalizeDatabase() {
	log.Println("Closing Database connection..")
	db.Finalize()
}
