package util

import "github.com/elbombardi/squrl/src/redirection_service/util"

func MockConfig() *util.Config {
	return &util.Config{
		Environment:    "test",
		DriverName:     "postgres",
		DBSource:       "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable",
		DBMaxIdleConns: 5,
		DBMaxOpenConns: 10,
		DBMaxIdleTime:  1,
		DBMaxLifeTime:  30,
		LogLevel:       "debug",
	}
}
