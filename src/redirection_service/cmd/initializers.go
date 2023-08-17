package main

import (
	"fmt"
	"log"

	"github.com/elbombardi/squrl/src/db"
	"github.com/elbombardi/squrl/src/redirection_service/routes"
	"github.com/elbombardi/squrl/src/redirection_service/util"
)

func initializeApp() (*routes.Server, error) {
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

	log.Println("Initializing App. Services..")
	return routes.NewServer(port, host, &routes.Routes{
		AccountRepository: store,
		URLRepository:     store,
		ClickRepository:   store,
	}), nil
}
