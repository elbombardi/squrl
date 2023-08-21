package main

import (
	"fmt"
	"log/slog"

	"github.com/elbombardi/squrl/src/api_service/api"
	"github.com/elbombardi/squrl/src/api_service/api/operations"
	"github.com/elbombardi/squrl/src/api_service/core"
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
	logger := util.NewLogger(&config)

	// Initializing db connection
	slog.Info("Initializing database connection..")
	store, err := db.GetStoreInstance(db.DBConf{
		DriverName:     config.DriverName,
		DataSourceName: config.DBSource,
		MaxIdleConns:   config.DBMaxIdleConns,
		MaxOpenConns:   config.DBMaxOpenConns,
		MaxIdleTime:    config.DBMaxIdleTime,
		MaxLifeTime:    config.DBMaxLifeTime,
	})
	if err != nil {
		slog.Error("Unable to initialize database connection", "Details", err)
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
		slog.Info("DB schema migration successful")
	}

	// Initializing API server
	swaggerSpec, err := loads.Embedded(api.SwaggerJSON, api.FlatSwaggerJSON)
	if err != nil {
		slog.Error("Error while initializing API server", "Details", err)
		return nil, err
	}
	adminAPI := operations.NewAdminAPI(swaggerSpec)
	adminAPI.Logger = func(s string, i ...interface{}) {
		slog.Info(fmt.Sprintf(s, i...))
	}
	server := api.NewServer(adminAPI)

	// Initializing services
	slog.Info("Initializing services..")
	authenticator := core.AuthenticationService{
		AccountRepository: store,
		Config:            &config,
		Logger:            logger.With("service", "AuthenticationService"),
	}
	accountsManager := core.AccountsService{
		AccountRepository: store,
		Config:            &config,
		Logger:            logger.With("service", "AccountsService"),
	}
	linksManager := core.LinksService{
		AccountRepository: store,
		LinkRepository:    store,
		Config:            &config,
		Logger:            logger.With("service", "LinksService"),
	}

	// Initializing endpoint handlers
	slog.Info("Initializing Handlers..")
	handlers := &handlers.Handlers{
		Authenticator:   &authenticator,
		AccountsManager: &accountsManager,
		LinksManager:    &linksManager,
		Config:          &config,
	}
	handlers.InstallHandlers(adminAPI)

	return server, nil
}
