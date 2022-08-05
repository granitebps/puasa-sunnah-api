package routes

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/granitebps/puasa-sunnah-api/docs"
	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/spf13/viper"
)

func InitRoutes(app *fiber.App) *fiber.App {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Puasa Sunnah API")
	})

	docs.SwaggerInfo.Host = viper.GetString("SWAGGER_HOST")
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	api := app.Group("api/v1")

	SourcesRoutes(api)
	CategoriesRoutes(api)
	TypesRoutes(api)
	FastingsRoutes(api)

	// Endpoint not found handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(http.StatusNotFound).JSON(helpers.FailedAPIResponse(
			"Endpoint not found",
			http.StatusNotFound,
		))
	})

	return app
}
