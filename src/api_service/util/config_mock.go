package util

func MockConfig() *Config {
	return &Config{
		Environment:        "test",
		DriverName:         "postgres",
		DBSource:           "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable",
		DBMaxIdleConns:     5,
		DBMaxOpenConns:     10,
		DBMaxIdleTime:      1,
		DBMaxLifeTime:      30,
		TokenSymmetricKey:  "test",
		AdminPassword:      "test",
		RedirectionBaseURL: "http://localhost:8080",
		LogLevel:           "debug",
	}
}
