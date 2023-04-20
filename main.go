package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/granitebps/puasa-sunnah-api/config"
	"github.com/granitebps/puasa-sunnah-api/docs"
	"github.com/granitebps/puasa-sunnah-api/routes"
	"github.com/spf13/viper"

	"github.com/gofiber/fiber/v2"
)

func LoadEnv() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
}

// @title Puasa Sunnah API
// @description This is a Puasa Sunnah API Docs
// @contact.name Granite Bagas
// @contact.email granitebagas28@gmail.com
// @license.name MIT
// @BasePath /
// @version 1.0
func main() {
	config.InitConfig(".env")

	docs.SwaggerInfo.Host = viper.GetString("SWAGGER_HOST")

	// Initialize Routes
	app := routes.InitRoutes()

	listenAndServe(app)
}

func listenAndServe(app *fiber.App) {
	PORT := viper.GetString("PORT")

	// Listen from a different goroutine
	go func() {
		if err := app.Listen(PORT); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")

	// Your cleanup tasks go here
	// db.Close()
	// redisConn.Close()
	fmt.Println("Fiber was successful shutdown.")
}
