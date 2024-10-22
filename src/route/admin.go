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
	source.Delete(":id", c.AdminController.DeleteSource)

	types := app.Group("types")
	types.Post("", c.AdminController.CreateType)
	types.Put(":id", c.AdminController.UpdateType)
	types.Delete(":id", c.AdminController.DeleteType)

	fasting := app.Group("fastings")
	fasting.Post("", c.AdminController.CreateFasting)
	fasting.Put(":id", c.AdminController.UpdateFasting)
}
