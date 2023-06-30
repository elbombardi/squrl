package db

import (
	"context"
	"database/sql"

	"github.com/elbombardi/squrl/util"
	_ "github.com/lib/pq"
)

type CustomersRepository interface {
	CheckApiKeyExists(ctx context.Context, apiKey string) (bool, error)
	CheckPrefixExists(ctx context.Context, prefix string) (bool, error)
	CheckUsernameExists(ctx context.Context, username string) (bool, error)
	GetCustomerByApiKey(ctx context.Context, apiKey string) (Customer, error)
	GetCustomerByPrefix(ctx context.Context, prefix string) (Customer, error)
	GetCustomerByUsername(ctx context.Context, username string) (Customer, error)
	InsertNewCustomer(ctx context.Context, arg InsertNewCustomerParams) error
	UpdateCustomerStatusByUsername(ctx context.Context, arg UpdateCustomerStatusByUsernameParams) error
}

type ShortURLsRepository interface {
	CheckShortUrlKeyExists(ctx context.Context, arg CheckShortUrlKeyExistsParams) (bool, error)
	GetShortURLByCustomerIDAndShortURLKey(ctx context.Context, arg GetShortURLByCustomerIDAndShortURLKeyParams) (ShortUrl, error)
	IncrementShortURLClickCount(ctx context.Context, id int32) error
	InsertNewShortURL(ctx context.Context, arg InsertNewShortURLParams) error
	SetShortURLFirstClickDate(ctx context.Context, arg SetShortURLFirstClickDateParams) error
	SetShortURLLastClickDate(ctx context.Context, arg SetShortURLLastClickDateParams) error
	UpdateShortURLLongURL(ctx context.Context, arg UpdateShortURLLongURLParams) error
	UpdateShortURLStatus(ctx context.Context, arg UpdateShortURLStatusParams) error
	UpdateShortURLTrackingStatus(ctx context.Context, arg UpdateShortURLTrackingStatusParams) error
}
type ClicksRepository interface {
	InsertNewClick(ctx context.Context, arg InsertNewClickParams) error
}

type Store interface {
	CustomersRepository
	ShortURLsRepository
	ClicksRepository
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

var dbInstance *sql.DB

func GetStoreInstance() (*SQLStore, error) {
	if dbInstance == nil {
		var err error
		dbInstance, err = sql.Open(util.ConfigDBDriver(), *util.ConfigDBSource())
		if err != nil {
			return nil, err
		}
		err = dbInstance.Ping()
		if err != nil {
			return nil, err
		}
		value, _ := util.ConfigDBMaxIdleConns()
		dbInstance.SetMaxIdleConns(value)
		value, _ = util.ConfigDBMaxOpenConns()
		dbInstance.SetMaxOpenConns(value)
		duration, _ := util.ConfigDBConnMaxIdleTime()
		dbInstance.SetConnMaxIdleTime(duration)
		duration, _ = util.ConfigDBConnMaxLifeTime()
		dbInstance.SetConnMaxLifetime(duration)
	}

	return &SQLStore{
		db:      dbInstance,
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
	tx, err := store.db.BeginTx(ctx, nil)
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
