package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/errors"
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
		return helpers.FailedAPIResponse(
			c,
			errors.WrapUserMessageAndCode(err, "", 0),
			nil,
		)
	}

	return helpers.SuccessAPIResponse(
		c,
		"Success",
		fiber.StatusOK,
		data,
	)
}
