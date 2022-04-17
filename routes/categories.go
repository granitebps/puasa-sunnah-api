package routes

import (
	"github.com/gofiber/fiber/v2"
	controlers "github.com/granitebps/puasa-sunnah-api/controllers"
)

func CategoriesRoutes(app fiber.Router) {
	api := app.Group("categories")

	api.Get("/", controlers.CategoriesIndex)
}
