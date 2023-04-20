package config

import (
	"fmt"
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
)

type Config struct {
	// Database   *Database
	// Log *Log
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
		// TODO: Use logrus
		log.Panic(err)
	}

	// log := NewLog(LOG_FOLDER)
	// newrelic := initNewrelic()
	// TODO: Parse log
	InitSentry()

	return &Config{
		// Database:   NewDb(),
		// Cache:      NewCache(newrelic.Application),
		// Log:        log,
		// HttpClient: NewHttpClient(log),
		// NewRelic:   newrelic,
	}
}

func errorConfig(err error) {
	fmt.Println("Error occured when setup config: ", err)
	sentry.CaptureException(err)
	panic(err)
}
