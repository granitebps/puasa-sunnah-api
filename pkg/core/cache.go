package core

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	storeRedis "github.com/gofiber/storage/redis/v3"
	"github.com/granitebps/puasa-sunnah-api/pkg/constants"
	nrredis "github.com/newrelic/go-agent/v3/integrations/nrredis-v9"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type Cache struct {
	Redis        *redis.Client
	RedisStorage *storeRedis.Storage
}

func SetupCache() *Cache {
	redisAddress := fmt.Sprintf("%s:%d", viper.GetString(constants.REDIS_HOST), viper.GetInt(constants.REDIS_PORT))
	redisPassword := viper.GetString(constants.REDIS_PASSWORD)
	redisDb := viper.GetInt(constants.REDIS_DB)

	redisConfig := &redis.Options{
		Addr:     redisAddress,
		Password: redisPassword,
		DB:       redisDb,
	}

	redisStorageConfig := storeRedis.Config{
		Addrs:    []string{redisAddress},
		Password: redisPassword,
		Database: redisDb,
	}

	if viper.GetString(constants.APP_ENV) != constants.LOCAL {
		//#nosec G402 -- Tested TLS version is 1.0
		redisConfig.TLSConfig = &tls.Config{}
		//#nosec G402 -- Tested TLS version is 1.0
		redisStorageConfig.TLSConfig = &tls.Config{}
	}

	return &Cache{
		Redis:        setupRedis(redisConfig),
		RedisStorage: setupRedisStorage(redisStorageConfig),
	}
}

func setupRedis(redisConfig *redis.Options) *redis.Client {
	rdb := redis.NewClient(redisConfig)

	rdb.AddHook(nrredis.NewHook(redisConfig))

	// Ping Redis Server
	ctx := context.Background()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		log.Panic(err)
	}

	return rdb
}

func setupRedisStorage(redisStorageConfig storeRedis.Config) *storeRedis.Storage {
	return storeRedis.New(redisStorageConfig)
}
