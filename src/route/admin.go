package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/src/controller"
)

func AdminRoutes(app fiber.Router, c *controller.ControllerStruct) {
	category := app.Group("categories")

	category.Post("", c.AdminController.CreateCategory)
	category.Put(":id", c.AdminController.UpdateCategory)

	source := app.Group("sources")
	source.Post("", c.AdminController.CreateSource)
	source.Put(":id", c.AdminController.UpdateSource)
}
