package scheduler

import (
	"context"

	"github.com/ansel1/merry/v2"
	"github.com/granitebps/puasa-sunnah-api/pkg/constants"
	"github.com/granitebps/puasa-sunnah-api/pkg/core"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/spf13/viper"
)

func SendHealthCheckSignal(c *core.Core) {
	// Initiate newrelic transaction
	txn := c.Newrelic.StartTransaction("SendHealthCheckSignal")
	defer txn.End()
	ctx := newrelic.NewContext(context.Background(), txn)

	// Send health check
	url := viper.GetString(constants.HEALTHCHECK_URL)
	_, err := c.Client.SetContext(ctx).Get(url)
	if err != nil {
		err = merry.Wrap(err)
		c.Log.Error(err)
	}
}
