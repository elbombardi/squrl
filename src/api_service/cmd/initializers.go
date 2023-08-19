package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/elbombardi/squrl/src/api_service/api"
	"github.com/elbombardi/squrl/src/api_service/api/operations"
	"github.com/elbombardi/squrl/src/api_service/handlers"
	"github.com/elbombardi/squrl/src/api_service/util"
	"github.com/elbombardi/squrl/src/db"
	"github.com/go-openapi/loads"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func initializeApp() (*api.Server, error) {
	// Loading configuration
	config, err := util.LoadConfig()
	if err != nil {
		slog.Error("Error while loading configuration", "Detail", err)
		return nil, err
	}

	// Initializing logger
	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: util.LogLevel(config.LogLevel),
	})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)
	util.LogConfig(&config)

	// Initializing API server
	swaggerSpec, err := loads.Embedded(api.SwaggerJSON, api.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}
	adminAPI := operations.NewAdminAPI(swaggerSpec)
	server := api.NewServer(adminAPI)

	// Initializing db connection
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
		return nil, fmt.Errorf("could not connect to the database")
	}

	// Checking DB schema migration
	slog.Info("Check for db schema changes..")
	driver, err := postgres.WithInstance(store.DB, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://../db/migration",
		"postgres", driver)
	if err != nil {
		slog.Error("Error preparing db migrations", "Details", err)
		return nil, err
	}
	err = m.Up()
	if err != nil {
		if err != migrate.ErrNoChange {
			slog.Error("Error executing db migrations", "Details", err)
			return nil, err
		}
		slog.Info("No db schema change")
	} else {
		slog.Info("DB Schema migration successful")
	}

	// Initializing endpoints handlers
	slog.Info("Initializing Handlers..")
	handlers := &handlers.Handlers{
		AccountRepository: store,
		URLRepository:     store,
		ClickRepository:   store,
		Config:            &config,
	}
	handlers.InstallHandlers(adminAPI)

	return server, nil
}
