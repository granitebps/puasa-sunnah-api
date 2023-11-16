package config

import (
	"github.com/getsentry/sentry-go"
	"github.com/granitebps/puasa-sunnah-api/pkg/constants"
	"github.com/spf13/viper"
)

func setupSentry() {
	debug := false
	appDebug := viper.GetBool(constants.APP_DEBUG)
	appEnv := viper.GetString(constants.APP_ENV)
	if appDebug && appEnv == constants.LOCAL {
		debug = true
	}

	dsn := viper.GetString(constants.SENTRY_DSN)
	_ = sentry.Init(sentry.ClientOptions{
		Dsn:              dsn,
		Debug:            debug,
		AttachStacktrace: true,
	})
}
