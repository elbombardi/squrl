package main

import (
	"fmt"
	"log"

	"github.com/elbombardi/squrl/api_service/handlers"
	db "github.com/elbombardi/squrl/db/sqlc"
	"github.com/elbombardi/squrl/util"
)

func initializeApp() (*handlers.Handlers, error) {
	err := util.LoadConfig()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	log.Println("Initializing Database connection..")
	store, err := db.GetStoreInstance()
	if err != nil {
		log.Println("Unable to initialize connection de database : ", err)
		return nil, err
	}
	if store == nil {
		log.Println("Could not connect to the database")
		return nil, fmt.Errorf("could not connect to the database")
	}

	log.Println("Initializing App. Services..")
	return &handlers.Handlers{
		CustomersRepository: store,
		ShortURLsRepository: store,
		ClicksRepository:    store,
	}, nil
}
