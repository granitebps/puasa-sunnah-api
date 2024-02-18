package controller

import (
	"github.com/ansel1/merry/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/pkg/core"
	"github.com/granitebps/puasa-sunnah-api/pkg/utils"
	"github.com/granitebps/puasa-sunnah-api/src/requests"
	"github.com/granitebps/puasa-sunnah-api/src/service"
)

type AdminController struct {
	AdminService *service.AdminService
	Core         *core.Core
}

func newAdminController(c *core.Core, AdminService *service.AdminService) *AdminController {
	return &AdminController{
		Core:         c,
		AdminService: AdminService,
	}
}

// Create Category godoc
// @Summary      Create category
// @Description  Create fasting category
// @Tags         Admin
// @Accept       json
// @Produce      json
// @param payload body requests.CreateCategoryRequest true "JSON payload"
// @Success      200  {object}   utils.JSONResponse{data=transformer.CategoryTransformer} "desc"
// @Failure      400  {object}  utils.JSONResponse
// @Router       /api/v1/admin/categories [post]
// @Security BasicAuth
func (c *AdminController) CreateCategory(ctx *fiber.Ctx) error {
	var req requests.CreateCategoryRequest
	errorField, err := c.Core.Validator.Validate(ctx, &req)
	if err != nil {
		err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusUnprocessableEntity))
		return utils.ReturnErrorResponse(ctx, err, errorField)
	}

	data, err := c.AdminService.CreateCategory(ctx.UserContext(), &req)
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(ctx, err, nil)
	}

	return utils.ReturnSuccessResponse(ctx, fiber.StatusOK, "Success create category", data)
}
