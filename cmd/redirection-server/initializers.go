package main

import (
	"fmt"
	"log"

	db "github.com/elbombardi/squrl/db/sqlc"
	"github.com/elbombardi/squrl/redirection_service"
	"github.com/elbombardi/squrl/redirection_service/routes"
	"github.com/elbombardi/squrl/util"
)

func initializeApp() (*routes.Routes, error) {
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

	log.Println("Starting Persistence Pool..")
	pool := redirection_service.NewPersistencePool(
		util.ConfigRedirectionPersistencePoolSize(),
		store,
		store,
	)
	pool.Start()

	log.Println("Initializing App. Services..")
	return &routes.Routes{
		CustomersRepository: store,
		ShortURLsRepository: store,
		PersistClick: func(shortUrl *db.ShortUrl, ipAddress string, userAgent string) {
			pool.AddJob(shortUrl, ipAddress, userAgent)
		},
	}, nil
}
