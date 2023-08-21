package db

import (
	"context"
	"database/sql"
	"time"

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

func GetStoreInstance(conf DBConf) (*SQLStore, error) {
	if dbInstance == nil {
		var err error
		dbInstance, err = sql.Open(conf.DriverName, conf.DataSourceName)
		if err != nil {
			return nil, err
		}
		err = dbInstance.Ping()
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

func Finalize() error {
	if dbInstance == nil {
		return nil
	}
	return dbInstance.Close()
}

func Transactional(ctx context.Context, db *sql.DB, fn func(queries *Queries) error) error {
	tx, err := db.BeginTx(ctx, nil)
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
