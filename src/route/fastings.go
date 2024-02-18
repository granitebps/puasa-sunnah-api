package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/src/controller"
)

func FastingsRoutes(app fiber.Router, c *controller.ControllerStruct) {
	api := app.Group("fastings")

	api.Get("/", c.FastingController.Index)
}
