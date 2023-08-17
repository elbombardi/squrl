package main

import (
	"fmt"
	"log"

	"github.com/elbombardi/squrl/src/api_service/handlers"
	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func initializeApp() (*handlers.Handlers, error) {
	config, err := util.LoadConfig()
	if err != nil {
		log.Print(err)
		return nil, err
	}

	log.Println("Initializing Database connection..")
	store, err := db.GetStoreInstance(db.DBConf{
		DriverName:     config.DriverName,
		DataSourceName: config.DBSource,
		MaxIdleConns:   config.DBMaxIdleConns,
		MaxOpenConns:   config.DBMaxOpenConns,
		MaxIdleTime:    config.DBMaxIdleTime,
		MaxLifeTime:    config.DBMaxLifeTime,
	})
	if err != nil {
		log.Println("Unable to initialize connection de database : ", err)
		return nil, err
	}
	if store == nil {
		log.Println("Could not connect to the database")
		return nil, fmt.Errorf("could not connect to the database")
	}

	log.Println("Check for database schema changes..")
	driver, err := postgres.WithInstance(store.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://../db/migration",
		"postgres", driver)
	if err != nil {
		log.Fatalln("ERROR While executing migrations : ", err)
	}
	m.Up()

	log.Println("Initializing App. Services..")
	return &handlers.Handlers{
		AccountRepository: store,
		URLRepository:     store,
		ClickRepository:   store,
		Config:            &config,
	}, nil
}
