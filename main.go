package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	config "github.com/granitebps/puasa-sunnah-api/configs"
	"github.com/granitebps/puasa-sunnah-api/pkg/constants"
	"github.com/granitebps/puasa-sunnah-api/pkg/core"
	"github.com/granitebps/puasa-sunnah-api/route"
	"github.com/granitebps/puasa-sunnah-api/src/middleware"
	"github.com/spf13/viper"

	"github.com/gofiber/fiber/v2"
)

// @title Puasa Sunnah API
// @description This is a Puasa Sunnah API Docs
// @contact.name Granite Bagas
// @contact.email granitebagas28@gmail.com
// @license.name MIT
// @BasePath /
// @version 1.0
func main() {
	// Load ENV and setup some config
	config.SetupConfig(".env")

	// Initiate Fiber
	app := fiber.New(config.FiberConfig())

	// Setup core package
	conf := core.SetupCore()

	// Setup middleware
	middleware.SetupMiddleware(app, conf)

	// Setup Dependency Injection
	contr := SetupDependencies(conf)

	// Setup route
	route.SetupRoute(app, contr)

	startServerWithGracefulShutdown(app)
}

func startServerWithGracefulShutdown(app *fiber.App) {
	PORT := viper.GetString(constants.APP_PORT)

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(PORT); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	<-c // This blocks the main thread until an interrupt is received
	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	log.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	log.Println("Fiber was successful shutdown.")
}
