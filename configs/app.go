package config

import (
	"log"

	"github.com/spf13/viper"
)

func SetupConfig(path string) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}

	setupSentry()
}
