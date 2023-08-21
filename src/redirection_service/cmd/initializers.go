package main

import (
	"fmt"
	"log/slog"

	"github.com/elbombardi/squrl/src/db"
	"github.com/elbombardi/squrl/src/redirection_service/core"
	"github.com/elbombardi/squrl/src/redirection_service/routes"
	"github.com/elbombardi/squrl/src/redirection_service/util"
)

func initializeApp() (*routes.Server, error) {
	// Loading configuration
	config, err := util.LoadConfig()
	if err != nil {
		slog.Error("Error while loading configuration", "Details", err)
		return nil, err
	}

	// Initializing logger
	logger := util.NewLogger(&config)

	// Initializing DB connection
	slog.Info("Initializing Database connection..")
	store, err := db.GetStoreInstance(db.DBConf{
		DriverName:     config.DriverName,
		DataSourceName: config.DBSource,
		MaxIdleConns:   config.DBMaxIdleConns,
		MaxOpenConns:   config.DBMaxOpenConns,
		MaxIdleTime:    config.DBMaxIdleTime,
		MaxLifeTime:    config.DBMaxLifeTime,
	})
	if err != nil {
		slog.Error("Unable to initialize connection de database", "Details", err)
		return nil, err
	}
	if store == nil {
		slog.Error("Could not connect to the database")
		return nil, fmt.Errorf("Could not connect to the database")
	}

	// Intializing services
	slog.Info("Initializing Redirection Service..")
	linksService := &core.LinksService{
		LinkRepository:    store,
		AccountRepository: store,
		ClickRepository:   store,
		Config:            &config,
		Logger:            logger.With("service", "LinksService"),
	}

	// Initalizing routes
	slog.Info("Initializing Redirection Server..")
	return routes.NewServer(port, host, &routes.Routes{
		LinksManager: linksService,
		Config:       &config,
	}), nil
}
