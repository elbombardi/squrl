package main

import (
	"log/slog"

	"github.com/elbombardi/squrl/src/db"
)

func finalizeApp() {

	slog.Info("Closing Database connection..")
	db.Finalize()
}
