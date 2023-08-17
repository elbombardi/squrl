package util

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
// The values are read by viper from a config file or environment variable.
type Config struct {
	Environment    string        `mapstructure:"ENVIRONMENT"`
	DriverName     string        `mapstructure:"DB_DRIVER"`
	DBSource       string        `mapstructure:"DB_SOURCE"`
	DBMaxIdleConns int           `mapstructure:"DB_MAX_IDLE_CONNS"`
	DBMaxOpenConns int           `mapstructure:"DB_MAX_OPEN_CONNS"`
	DBMaxIdleTime  time.Duration `mapstructure:"DB_MAX_IDLE_TIME"`
	DBMaxLifeTime  time.Duration `mapstructure:"DB_MAX_LIFE_TIME"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.SetDefault("ENVIRONMENT", "prod")
	viper.SetDefault("DB_DRIVER", "postgres")
	viper.SetDefault("DB_MAX_IDLE_CONNS", 5)
	viper.SetDefault("DB_MAX_OPEN_CONNS", 10)
	viper.SetDefault("DB_MAX_IDLE_TIME", 1*time.Second)
	viper.SetDefault("DB_MAX_LIFE_TIME", 30*time.Second)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	requiredConfig := []string{
		"DB_SOURCE",
	}
	for _, key := range requiredConfig {
		v := viper.Get(key)
		if v == nil {
			return config, fmt.Errorf("Missing required configuration: %s", key)
		}
	}

	if config.Environment == "dev" {
		log.Println("Configuration loaded : ")
		for _, key := range viper.AllKeys() {
			log.Printf("\t* %s = %v\n", key, viper.Get(key))
		}
	}

	return
}
