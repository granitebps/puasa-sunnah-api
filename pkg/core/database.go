package core

import (
	"database/sql"
	"log"
	"time"

	"github.com/go-gorm/caches/v4"
	"github.com/granitebps/puasa-sunnah-api/pkg/constants"
	"github.com/granitebps/puasa-sunnah-api/src/model"
	_ "github.com/newrelic/go-agent/v3/integrations/nrmysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Db *gorm.DB
}

func SetupDb() *Database {
	return &Database{
		Db: SetupMySql(),
	}
}

func SetupMySql() *gorm.DB {
	host := viper.GetString(constants.DB_HOST)
	port := viper.GetString(constants.DB_PORT)
	user := viper.GetString(constants.DB_USER)
	pass := viper.GetString(constants.DB_PASS)
	name := viper.GetString(constants.DB_NAME)

	var dsn = ""
	if user != "" {
		dsn += user
	}
	if pass != "" {
		dsn = dsn + ":" + pass
	}
	if host != "" {
		dsn = dsn + "@tcp(" + host
		if port != "" {
			dsn = dsn + ":" + port + ")"
		} else {
			dsn += ")"
		}
	}
	if name != "" {
		dsn = dsn + "/" + name
	}
	dsn += "?parseTime=true&interpolateParams=true"

	db, err := sql.Open("nrmysql", dsn)
	if err != nil {
		log.Panic(err)
	}

	mode := logger.Silent
	if viper.GetString(constants.APP_ENV) == constants.LOCAL {
		mode = logger.Info
	}

	database, err := gorm.Open(mysql.New(mysql.Config{
		Conn: db,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(mode),
		PrepareStmt:            true,
	})
	if err != nil {
		log.Panic(err)
	}

	cachesPlugin := &caches.Caches{Conf: &caches.Config{
		Easer: true,
	}}

	database.Set("gorm:auto_preload", true)
	database.Session(&gorm.Session{
		AllowGlobalUpdate:    true,
		FullSaveAssociations: false,
	})

	sDatabase, _ := database.DB()
	sDatabase.SetMaxOpenConns(20)
	sDatabase.SetMaxIdleConns(5)
	sDatabase.SetConnMaxIdleTime(1 * time.Minute)
	sDatabase.SetConnMaxLifetime(1 * time.Minute)

	// Cache
	err = database.Use(cachesPlugin)
	if err != nil {
		log.Panic(err)
	}

	// Auto migrate
	err = database.AutoMigrate(&model.Category{})
	if err != nil {
		log.Panic(err)
	}
	err = database.AutoMigrate(&model.Type{})
	if err != nil {
		log.Panic(err)
	}
	err = database.AutoMigrate(&model.Source{})
	if err != nil {
		log.Panic(err)
	}
	err = database.AutoMigrate(&model.Fasting{})
	if err != nil {
		log.Panic(err)
	}

	return database
}
