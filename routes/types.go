package routes

import (
	"github.com/gofiber/fiber/v2"
	controlers "github.com/granitebps/puasa-sunnah-api/controllers"
)

func TypesRoutes(app fiber.Router) {
	api := app.Group("types")

	api.Get("/", controlers.TypesIndex)
}
