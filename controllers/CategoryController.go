package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/errors"
	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/services"
)

type CategoryController struct {
	CategoryService *services.CategoryService
}

func newCategoryController(categoryService *services.CategoryService) *CategoryController {
	return &CategoryController{
		CategoryService: categoryService,
	}
}

// ListCategory godoc
// @Summary      List Categories
// @Description  Get list of categories
// @Tags         Categories
// @Accept       json
// @Produce      json
// @Success      200  {object}   helpers.SuccessResponse{data=[]types.Category} "desc"
// @Failure      400  {object}  helpers.FailedResponse
// @Router       /api/v1/categories [get]
func (c *CategoryController) Index(ctx *fiber.Ctx) error {
	data, err := c.CategoryService.GetAll()
	if err != nil {
		return helpers.FailedAPIResponse(
			ctx,
			errors.WrapUserMessageAndCode(err, "", 0),
			nil,
		)
	}

	return helpers.SuccessAPIResponse(
		ctx,
		"Success",
		fiber.StatusOK,
		data,
	)
}
