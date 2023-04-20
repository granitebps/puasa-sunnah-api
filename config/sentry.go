package config

import (
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
)

// TODO: Parse log
func InitSentry() {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: viper.GetString("SENTRY_DSN"),
	}); err != nil {
		// log.Logger.Panic(e.WrapError(err))
		log.Panic(err)
	}
}
