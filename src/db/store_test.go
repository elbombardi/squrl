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
		MaxIdleConns:   10,
		MaxOpenConns:   10,
		MaxIdleTime:    10,
		MaxLifeTime:    10,
	})

	if err != nil {
		panic(err)
	}

	code := m.Run()

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
