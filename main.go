package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/granitebps/puasa-sunnah-api/configs"
	"github.com/granitebps/puasa-sunnah-api/controllers"
	"github.com/granitebps/puasa-sunnah-api/docs"
	"github.com/granitebps/puasa-sunnah-api/repositories"
	"github.com/granitebps/puasa-sunnah-api/routes"
	"github.com/granitebps/puasa-sunnah-api/services"
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
	configApp := configs.InitConfig(".env")
	configApp.Log.Logger.Info("Puasa Sunnah API")

	docs.SwaggerInfo.Host = viper.GetString("SWAGGER_HOST")

	// Init repo
	sourceRepo := repositories.NewSourceRepository(configApp)
	typesRepo := repositories.NewTypesRepository(configApp)
	categoryRepo := repositories.NewCategoryRepository(configApp)
	fastingRepo := repositories.NewFastingRepository(configApp)

	// Init service
	sourceService := services.NewSourceService(sourceRepo)
	typesService := services.NewTypesService(typesRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	fastingService := services.NewFastingService(fastingRepo, categoryRepo, typesRepo)

	controller := controllers.NewController(sourceService, typesService, categoryService, fastingService)

	// Initialize Routes
	app := routes.InitRoutes(configApp.Log, controller)

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
