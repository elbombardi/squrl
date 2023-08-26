package db

import (
	"os"
	"testing"
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
		MaxIdleConns:   1,
		MaxOpenConns:   20,
		MaxIdleTime:    1,
		MaxLifeTime:    1,
	})

	if err != nil {
		panic(err)
	}

	code := m.Run()

	testStore.DB.Close()

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
