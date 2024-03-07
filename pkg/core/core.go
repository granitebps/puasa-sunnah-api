package core

import (
	"github.com/go-resty/resty/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

type Core struct {
	Newrelic  *newrelic.Application
	Log       *logrus.Logger
	Client    *resty.Request
	Database  *Database
	Validator *AppValidator
}

func SetupCore() *Core {
	nr := SetupNewrelicApp()
	log := SetupLog()
	client := SetupResty()
	db := SetupDb()
	v := SetupValidator()

	return &Core{
		Newrelic:  nr,
		Log:       log,
		Client:    client,
		Database:  db,
		Validator: v,
	}
}
