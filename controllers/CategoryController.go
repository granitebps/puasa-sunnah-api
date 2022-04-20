package controlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/services"
)

// ListCategory godoc
// @Summary      List Categories
// @Description  Get list of categories
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Success      200  {object}   helpers.SuccessResponse{data=[]types.Category} "desc"
// @Failure      400  {object}  helpers.FailedResponse
// @Router       /api/v1/categories [get]
func CategoriesIndex(c *fiber.Ctx) error {
	data, err := services.CategoriesGetAll()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(helpers.FailedAPIResponse(
			err.Error(),
			http.StatusBadRequest,
		))
	}

	return c.Status(http.StatusOK).JSON(helpers.SuccessAPIResponse(
		"Success",
		http.StatusOK,
		data,
	))
}
