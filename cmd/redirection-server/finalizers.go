package main

import (
	"log"

	db "github.com/elbombardi/squrl/db/sqlc"
	"github.com/elbombardi/squrl/redirection_service"
)

func finalizeApp() {
	defer finalizeDatabase()
	finalizePersistencePool()
}

func finalizeDatabase() {
	log.Println("Closing Database connection..")
	db.Finalize()
}

func finalizePersistencePool() {
	redirection_service.FinalizePersistencePool()
}
