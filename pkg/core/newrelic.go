package core

import (
	"log"

	"github.com/gofiber/contrib/fibernewrelic"
	"github.com/granitebps/puasa-sunnah-api/pkg/constants"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/spf13/viper"
)

func SetupNewrelicApp() *newrelic.Application {
	app, err := newrelic.NewApplication(
		newrelic.ConfigEnabled(viper.GetBool(constants.NEWRELIC_ENABLED)),
		newrelic.ConfigAppName(viper.GetString(constants.NEWRELIC_APP_NAME)),
		newrelic.ConfigLicense(viper.GetString(constants.NEWRELIC_LICENSE)),
		// newrelic.ConfigDebugLogger(os.Stdout),
		// newrelic.ConfigAppLogForwardingEnabled(true),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		log.Panic(err)
	}

	return app
}

func SetupNewrelicFiber(a *newrelic.Application) fibernewrelic.Config {
	return fibernewrelic.Config{
		Application: a,
	}
}
