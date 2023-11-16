package controller

import (
	"github.com/ansel1/merry"
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/pkg/utils"
	"github.com/granitebps/puasa-sunnah-api/src/service"
)

type CategoryController struct {
	CategoryService *service.CategoryService
}

func newCategoryController(categoryService *service.CategoryService) *CategoryController {
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
// @Success      200  {object}   utils.JSONResponse{data=[]types.Category} "desc"
// @Failure      400  {object}  utils.JSONResponse
// @Router       /api/v1/categories [get]
func (c *CategoryController) Index(ctx *fiber.Ctx) error {
	data, err := c.CategoryService.GetAll()
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(ctx, err, nil)
	}

	return utils.ReturnSuccessResponse(ctx, fiber.StatusOK, "Success", data)
}
