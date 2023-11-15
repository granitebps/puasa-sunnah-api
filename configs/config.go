package configs

import (
	"log"

	"github.com/granitebps/puasa-sunnah-api/pkg/constants"
	"github.com/spf13/viper"
)

type Config struct {
	// Database   *Database
	Log *Log
	// NewRelic   *NewRelic
	// HttpClient *HttpClient
	// Cache      *Cache
}

func LoadConfig(file string) error {
	viper.SetConfigFile(file)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	return nil
}

func InitConfig(path string) *Config {
	err := LoadConfig(path)
	if err != nil {
		log.Panic(err)
	}

	logger := NewLog(constants.LOG_FOLDER)
	// newrelic := initNewrelic()
	InitSentry(logger)

	return &Config{
		// Database:   NewDb(),
		// Cache:      NewCache(newrelic.Application),
		Log: logger,
		// HttpClient: NewHttpClient(log),
		// NewRelic:   newrelic,
	}
}
