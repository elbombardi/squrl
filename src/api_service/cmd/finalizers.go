package main

import (
	"log"

	"github.com/elbombardi/squrl/src/db"
)

func finalizeApp() {
	defer finalizeDatabase()
}

func finalizeDatabase() {
	log.Println("Closing Database connection..")
	db.Finalize()
}
