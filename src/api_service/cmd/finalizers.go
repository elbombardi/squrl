package main

import (
	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
)

func finalizeApp() {
	defer finalizeDatabase()
}

func finalizeDatabase() {
	util.Info("Closing Database connection..")
	db.Finalize()
}
