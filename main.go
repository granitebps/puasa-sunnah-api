package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/granitebps/puasa-sunnah-api/docs"
	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/routes"
	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/gofiber/swagger"
)

// @title Puasa Sunnah API
// @description This is a Puasa Sunnah API Docs
// @contact.name Granite Bagas
// @contact.email granitebagas28@gmail.com
// @license.name MIT
// @BasePath /
func main() {
	// Load ENV
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}
	PORT := os.Getenv("PORT")

	app := fiber.New()

	// Middleware
	// Define file to logs
	now := time.Now().Format("2006-02-01")
	logFileName := "./logs/" + now + ".log"
	file, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	// Set config for logger
	loggerConfig := logger.Config{
		Output: file, // add file to save output
	}
	app.Use(logger.New(loggerConfig))
	app.Use(limiter.New(limiter.Config{
		Max: 60,
	}))
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Puasa Sunnah API")
	})

	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	api := app.Group("api/v1")

	routes.SourcesRoutes(api)
	routes.CategoriesRoutes(api)
	routes.TypesRoutes(api)
	routes.FastingsRoutes(api)

	// Endpoint not found handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON(helpers.FailedAPIResponse(
			"Endpoint not found",
			http.StatusNotFound,
		))
	})

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
