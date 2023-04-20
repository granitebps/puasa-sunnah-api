package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/granitebps/puasa-sunnah-api/configs"
	"github.com/granitebps/puasa-sunnah-api/constants"
	e "github.com/granitebps/puasa-sunnah-api/errors"
	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/middleware"
)

func InitRoutes(log *configs.Log) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			log.Logger.Error(err)

			var code int
			var msg string
			if helpers.IsProduction() {
				code = fiber.StatusInternalServerError
				msg = constants.INTERNAL_SERVER_ERROR
			} else {
				var e *fiber.Error
				if errors.As(err, &e) {
					code = e.Code
					msg = e.Message
				} else {
					msg = err.Error()
				}
			}

			return helpers.FailedAPIResponse(
				ctx,
				e.WrapUserMessageAndCode(err, msg, code),
				nil,
			)
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
		return helpers.FailedAPIResponse(
			c,
			e.WrapUserMessageAndCode(e.ErrEndpointNotFound, "", fiber.StatusNotFound),
			nil,
		)
	})

	return app
}
