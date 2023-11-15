package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/src/controller"
)

func SourcesRoutes(app fiber.Router, c *controller.ControllerStruct) {
	api := app.Group("sources")

	api.Get("/", c.SourceController.Index)
}
