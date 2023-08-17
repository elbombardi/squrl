package db

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type AccountRepository interface {
	CheckApiKeyExists(ctx context.Context, apiKey string) (bool, error)
	CheckPrefixExists(ctx context.Context, prefix string) (bool, error)
	CheckUsernameExists(ctx context.Context, username string) (bool, error)
	GetAccountByApiKey(ctx context.Context, apiKey string) (Account, error)
	GetAccountByPrefix(ctx context.Context, prefix string) (Account, error)
	GetAccountByUsername(ctx context.Context, username string) (Account, error)
	InsertNewAccount(ctx context.Context, arg InsertNewAccountParams) error
	UpdateAccountStatusByUsername(ctx context.Context, arg UpdateAccountStatusByUsernameParams) error
}

type URLRepository interface {
	CheckShortUrlKeyExists(ctx context.Context, arg CheckShortUrlKeyExistsParams) (bool, error)
	GetURLByAccountIDAndShortURLKey(ctx context.Context, arg GetURLByAccountIDAndShortURLKeyParams) (Url, error)
	InsertNewURL(ctx context.Context, arg InsertNewURLParams) error
	UpdateLongURL(ctx context.Context, arg UpdateLongURLParams) error
	UpdateURLStatus(ctx context.Context, arg UpdateURLStatusParams) error
	UpdateURLTrackingStatus(ctx context.Context, arg UpdateURLTrackingStatusParams) error
}
type ClickRepository interface {
	InsertNewClick(ctx context.Context, arg InsertNewClickParams) error
}

type Store interface {
	AccountRepository
	URLRepository
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
		Queries: &Queries{db: dbInstance},
	}, nil
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
