package routes

import (
	"github.com/gofiber/fiber/v2"
	controlers "github.com/granitebps/puasa-sunnah-api/controllers"
)

func SourcesRoutes(app fiber.Router, c *controlers.Controller) {
	api := app.Group("sources")

	api.Get("/", c.SourceController.Index)
}
