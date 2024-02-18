package route

import (
	"errors"
	"fmt"

	"github.com/ansel1/merry/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/granitebps/puasa-sunnah-api/pkg/constants"
	"github.com/granitebps/puasa-sunnah-api/pkg/utils"
	"github.com/granitebps/puasa-sunnah-api/src/controller"
	"github.com/spf13/viper"
)

func SetupRoute(a *fiber.App, c *controller.ControllerStruct) {
	a.Get("", func(ctx *fiber.Ctx) error {
		return utils.ReturnSuccessResponse(ctx, fiber.StatusOK, fmt.Sprintf("%s API", viper.GetString(constants.APP_NAME)), nil)
	})

	route := a.Group("api")

	// V1 Route
	v1Route(route, c)
}

func v1Route(route fiber.Router, c *controller.ControllerStruct) {
	v1 := route.Group("v1")

	SourcesRoutes(v1, c)
	CategoriesRoutes(v1, c)
	TypesRoutes(v1, c)
	FastingsRoutes(v1, c)

	// Admin route
	adminRoute(v1, c)
}

func adminRoute(route fiber.Router, c *controller.ControllerStruct) {
	admin := route.Group("admin")

	// Add basic auth
	admin.Use(basicauth.New(basicauth.Config{
		Users: map[string]string{
			viper.GetString(constants.ADMIN_USERNAME): viper.GetString(constants.ADMIN_PASSWORD),
		},
		Unauthorized: func(c *fiber.Ctx) error {
			err := errors.New("Unauthorized")
			err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusUnauthorized), merry.WithUserMessage("Unauthorized"))

			return utils.ReturnErrorResponse(c, err, nil)
		},
	}))

	AdminRoutes(admin, c)
}
