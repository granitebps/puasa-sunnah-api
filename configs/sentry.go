package configs

import (
	"github.com/getsentry/sentry-go"
	"github.com/spf13/viper"
)

func InitSentry(log *Log) {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: viper.GetString("SENTRY_DSN"),
	}); err != nil {
		log.Logger.Panic(err)
	}
}
