package routes

import (
	"github.com/gofiber/fiber/v2"
	controlers "github.com/granitebps/puasa-sunnah-api/controllers"
)

func FastingsRoutes(app fiber.Router) {
	api := app.Group("fastings")

	api.Get("/", controlers.FastingsIndex)
}
