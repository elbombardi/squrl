package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"

	_ "github.com/lib/pq"
)

type AccountRepository interface {
	CheckPrefixExists(ctx context.Context, prefix string) (bool, error)
	CheckUsernameExists(ctx context.Context, username string) (bool, error)
	GetAccountByPrefix(ctx context.Context, prefix string) (Account, error)
	GetAccountByUsername(ctx context.Context, username string) (Account, error)
	InsertNewAccount(ctx context.Context, arg InsertNewAccountParams) error
	UpdateAccountStatusByUsername(ctx context.Context, arg UpdateAccountStatusByUsernameParams) error
}

type LinkRepository interface {
	CheckShortUrlKeyExists(ctx context.Context, arg CheckShortUrlKeyExistsParams) (bool, error)
	GetLinkByAccountIDAndShortURLKey(ctx context.Context, arg GetLinkByAccountIDAndShortURLKeyParams) (Link, error)
	InsertNewLink(ctx context.Context, arg InsertNewLinkParams) error
	UpdateLinkLongURL(ctx context.Context, arg UpdateLinkLongURLParams) error
	UpdateLinkStatus(ctx context.Context, arg UpdateLinkStatusParams) error
	UpdateLinkTrackingStatus(ctx context.Context, arg UpdateLinkTrackingStatusParams) error
}
type ClickRepository interface {
	InsertNewClick(ctx context.Context, arg InsertNewClickParams) error
}

type Store interface {
	AccountRepository
	LinkRepository
	ClickRepository
	Transactional(ctx context.Context, fn func(queries *Queries) error) error
}

type SQLStore struct {
	*Queries
	DB *sql.DB
}

var dbInstance *sql.DB

type DBConf struct {
	DriverName     string
	DataSourceName string
	MaxIdleConns   int
	MaxOpenConns   int
	MaxIdleTime    time.Duration
	MaxLifeTime    time.Duration
}

// GetStoreInstance returns a singleton instance of the SQLStore
// If the instance is not initialized, it will initialize it
// with the provided configuration
// If the instance is already initialized, it will return the existing instance
func GetStoreInstance(conf DBConf) (*SQLStore, error) {

	if dbInstance == nil {
		var err error

		dbInstance, err = sql.Open(conf.DriverName, conf.DataSourceName)
		if err != nil {
			return nil, err
		}

		retries := 5
		for retries > 0 {
			retries--
			err = dbInstance.Ping()
			if err != nil {
				slog.Error("Error while pinging the database. Retrying in 5 seconds", "Details", err)
				time.Sleep(5 * time.Second)
			}
		}
		if err != nil {
			return nil, err
		}

		dbInstance.SetMaxIdleConns(conf.MaxIdleConns)
		dbInstance.SetMaxOpenConns(conf.MaxOpenConns)
		dbInstance.SetConnMaxIdleTime(conf.MaxIdleTime)
		dbInstance.SetConnMaxLifetime(conf.MaxLifeTime)
	}

	return &SQLStore{
		DB:      dbInstance,
		Queries: New(dbInstance),
	}, nil
}

// MigrateIfNeeded checks if the database schema needs to be migrated
// and performs the migration if needed.
// schemaURL is the path to the directory containing the migration files
// e.g. file://../db/migration
//
// Returns true if the migration was performed, false if no migration was needed
func MigrateIfNeeded(schemaURL string) (bool, error) {
	driver, err := postgres.WithInstance(dbInstance, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		schemaURL,
		"postgres", driver)
	if err != nil {
		return false, fmt.Errorf("Error preparing db migrations (%v)", err)
	}
	err = m.Up()
	if err != nil {
		if err != migrate.ErrNoChange {
			return false, fmt.Errorf("error executing db migrations (%v)", err)
		}
		// No Changes
		return false, nil
	}

	// Migration successful
	return true, nil
}

// DropAll drops all tables in the database
// schemaURL is the path to the directory containing the migration files
// e.g. file://../db/migration
//
// Returns an error if the drop failed
func DropAll(schemaURL string) error {
	driver, err := postgres.WithInstance(dbInstance, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		schemaURL,
		"postgres", driver)
	if err != nil {
		return fmt.Errorf("error dropping database (%v)", err)
	}
	err = m.Down()
	return err
}

func Finalize() error {
	if dbInstance == nil {
		return nil
	}
	return dbInstance.Close()
}

func (store *SQLStore) Transactional(ctx context.Context, fn func(queries *Queries) error) error {
	tx, err := store.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	return tx.Commit()
}
