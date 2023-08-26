package db

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
)

var testStore *SQLStore

func TestMain(m *testing.M) {
	var err error
	testStore, err = GetStoreInstance(DBConf{
		DriverName:     dbDriver,
		DataSourceName: dbSource,
		MaxIdleConns:   10,
		MaxOpenConns:   30,
		MaxIdleTime:    1,
		MaxLifeTime:    1,
	})
	if err != nil {
		panic(err)
	}

	code := m.Run()

	Finalize()

	os.Exit(code)
}

func setup() {
	dbSchemaDir := "file://./migration"
	_, err := MigrateIfNeeded(dbSchemaDir)
	if err != nil {
		panic(err)
	}
}

func teardown() {
	dbSchemaDir := "file://./migration"
	err := DropAll(dbSchemaDir)
	if err != nil {
		panic(err)
	}
}

func TestTransactional(t *testing.T) {
	setup()
	defer teardown()
	ctx := context.Background()
	err := testStore.Transactional(ctx, func(queries *Queries) error {
		return testStore.InsertNewAccount(ctx, InsertNewAccountParams{
			Prefix:         "tst",
			Username:       "username",
			Email:          "email@gmail.com",
			HashedPassword: "$2a$1",
		})
	})
	require.NoError(t, err, "Error should be nil")
}
