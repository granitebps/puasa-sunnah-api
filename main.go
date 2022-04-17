package main

import (
	"github.com/granitebps/puasa-sunnah-api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(limiter.New(limiter.Config{
		Max: 60,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Puasa Sunnah API")
	})

	api := app.Group("api/v1")

	routes.SourcesRoutes(api)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":3000")
}
