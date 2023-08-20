package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/elbombardi/squrl/src/db"
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
	logHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: util.LogLevel(config.LogLevel),
	})
	logger := slog.New(logHandler)
	logger = logger.With(
		slog.Group("program_info",
			slog.Int("pid", os.Getpid()),
			slog.String("component", "squrl.RedirectionService"),
			slog.String("version", util.VERSION),
			slog.String("environment", config.Environment),
		),
	)
	slog.SetDefault(logger)
	util.LogConfig(&config)

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

	// Initalizing server
	slog.Info("Initializing Redirection Server..")
	return routes.NewServer(port, host, &routes.Routes{
		AccountRepository: store,
		URLRepository:     store,
		ClickRepository:   store,
	}), nil
}
