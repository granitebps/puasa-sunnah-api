package core

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/granitebps/puasa-sunnah-api/pkg/constants"
	nrredis "github.com/newrelic/go-agent/v3/integrations/nrredis-v9"
	memory "github.com/patrickmn/go-cache"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type Cache struct {
	Redis  *redis.Client
	Memory *memory.Cache
}

func SetupCache() *Cache {

	return &Cache{
		Redis:  setupRedis(),
		Memory: setupInMemory(),
	}
}

func setupRedis() *redis.Client {
	opts := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString(constants.REDIS_HOST), viper.GetInt(constants.REDIS_PORT)),
		Password: viper.GetString(constants.REDIS_PASSWORD),
		DB:       viper.GetInt(constants.REDIS_DB),
	}

	if viper.GetString(constants.APP_ENV) != constants.LOCAL {
		//#nosec G402 -- Tested TLS version is 1.0
		opts.TLSConfig = &tls.Config{}
	}

	rdb := redis.NewClient(opts)

	rdb.AddHook(nrredis.NewHook(opts))

	// Ping Redis Server
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		log.Panic(err)
	}

	return rdb
}

func setupInMemory() *memory.Cache {
	c := memory.New(memory.NoExpiration, time.Minute*10)

	return c
}
