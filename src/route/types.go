package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/src/controller"
)

func TypesRoutes(app fiber.Router, c *controller.ControllerStruct) {
	api := app.Group("types")

	api.Get("/", c.TypesController.Index)
}
