package routes

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/middleware"
)

func InitRoutes() *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError
			msg := "Internal Server Error"

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
				msg = e.Message
			}

			return ctx.Status(code).JSON(helpers.FailedAPIResponse(
				msg,
				code,
			))
		},
	})

	// Initialize Middlewares
	middleware.InitMiddleware(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Puasa Sunnah API")
	})

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
