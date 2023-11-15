package core

import (
	"github.com/go-resty/resty/v2"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/sirupsen/logrus"
)

type Core struct {
	Newrelic  *newrelic.Application
	Log       *logrus.Logger
	Cache     *Cache
	Client    *resty.Request
	Database  *Database
	Validator *AppValidator
}

func SetupCore() *Core {
	nr := SetupNewrelicApp()
	log := SetupLog()
	cache := SetupCache()
	client := SetupResty()
	db := SetupDb()
	v := SetupValidator()

	return &Core{
		Newrelic:  nr,
		Log:       log,
		Cache:     cache,
		Client:    client,
		Database:  db,
		Validator: v,
	}
}
