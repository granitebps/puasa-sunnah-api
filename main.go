package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/granitebps/puasa-sunnah-api/middleware"
	"github.com/granitebps/puasa-sunnah-api/routes"
	"github.com/joho/godotenv"

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
	// Load ENV
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}
	PORT := os.Getenv("PORT")

	app := fiber.New()

	// Initialize Middlewares
	middleware.InitMiddleware(app)

	// Initialize Routes
	routes.InitRoutes(app)

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
