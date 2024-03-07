package controller

import (
	"strconv"

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

func newAdminController(c *core.Core, adminService *service.AdminService) *AdminController {
	return &AdminController{
		Core:         c,
		AdminService: adminService,
	}
}

// Create Category godoc
// @Summary      Create category
// @Description  Create fasting category
// @Tags         Admin
// @Accept       json
// @Produce      json
// @param payload body requests.CategoryRequest true "JSON payload"
// @Success      200  {object}   utils.JSONResponse{data=transformer.CategoryTransformer} "desc"
// @Failure      400  {object}  utils.JSONResponse
// @Router       /api/v1/admin/categories [post]
// @Security BasicAuth
func (c *AdminController) CreateCategory(ctx *fiber.Ctx) error {
	var req requests.CategoryRequest
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

// Update Category	godoc
// @Summary      	Update category
// @Description  	Update fasting category
// @Tags         	Admin
// @Accept       	json
// @Produce      	json
// @param 		 	payload body requests.CategoryRequest true "JSON payload"
// @Param 			id path int true "Category ID"
// @Success      	200  {object}  utils.JSONResponse{data=transformer.CategoryTransformer} "desc"
// @Failure      	400  {object}  utils.JSONResponse
// @Router       	/api/v1/admin/categories/:id [put]
// @Security 		BasicAuth
func (c *AdminController) UpdateCategory(ctx *fiber.Ctx) error {
	var req requests.CategoryRequest

	errorField, err := c.Core.Validator.Validate(ctx, &req)
	if err != nil {
		err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusUnprocessableEntity))
		return utils.ReturnErrorResponse(ctx, err, errorField)
	}

	idString := ctx.Params("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(ctx, err, nil)
	}

	data, err := c.AdminService.UpdateCategory(ctx.UserContext(), uint(id), &req)
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(ctx, err, nil)
	}

	return utils.ReturnSuccessResponse(ctx, fiber.StatusOK, "Success update category", data)
}

// Create Source	godoc
// @Summary      	Create source
// @Description  	Create fasting source
// @Tags         	Admin
// @Accept       	json
// @Produce      	json
// @param 			payload body requests.SourceRequest true "JSON payload"
// @Success      	200  {object}   utils.JSONResponse{data=transformer.SourceTransformer} "desc"
// @Failure      	400  {object}  utils.JSONResponse
// @Router       	/api/v1/admin/sources [post]
// @Security 		BasicAuth
func (c *AdminController) CreateSource(ctx *fiber.Ctx) error {
	var req requests.SourceRequest
	errorField, err := c.Core.Validator.Validate(ctx, &req)
	if err != nil {
		err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusUnprocessableEntity))
		return utils.ReturnErrorResponse(ctx, err, errorField)
	}

	data, err := c.AdminService.CreateSource(ctx.UserContext(), &req)
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(ctx, err, nil)
	}

	return utils.ReturnSuccessResponse(ctx, fiber.StatusOK, "Success create source", data)
}

// Update Source	godoc
// @Summary      	Update source
// @Description  	Update fasting source
// @Tags         	Admin
// @Accept       	json
// @Produce      	json
// @param 		 	payload body requests.SourceRequest true "JSON payload"
// @Param 			id path int true "Source ID"
// @Success      	200  {object}  utils.JSONResponse{data=transformer.SourceTransformer} "desc"
// @Failure      	400  {object}  utils.JSONResponse
// @Router       	/api/v1/admin/sources/:id [put]
// @Security 		BasicAuth
func (c *AdminController) UpdateSource(ctx *fiber.Ctx) error {
	var req requests.SourceRequest

	errorField, err := c.Core.Validator.Validate(ctx, &req)
	if err != nil {
		err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusUnprocessableEntity))
		return utils.ReturnErrorResponse(ctx, err, errorField)
	}

	idString := ctx.Params("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(ctx, err, nil)
	}

	data, err := c.AdminService.UpdateSource(ctx.UserContext(), uint(id), &req)
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(ctx, err, nil)
	}

	return utils.ReturnSuccessResponse(ctx, fiber.StatusOK, "Success update source", data)
}

// Create Type	godoc
// @Summary      	Create type
// @Description  	Create fasting type
// @Tags         	Admin
// @Accept       	json
// @Produce      	json
// @param 			payload body requests.TypeRequest true "JSON payload"
// @Success      	200  {object}   utils.JSONResponse{data=transformer.TypeTransformer} "desc"
// @Failure      	400  {object}  utils.JSONResponse
// @Router       	/api/v1/admin/types [post]
// @Security 		BasicAuth
func (c *AdminController) CreateType(ctx *fiber.Ctx) error {
	var req requests.TypeRequest
	errorField, err := c.Core.Validator.Validate(ctx, &req)
	if err != nil {
		err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusUnprocessableEntity))
		return utils.ReturnErrorResponse(ctx, err, errorField)
	}

	data, err := c.AdminService.CreateType(ctx.UserContext(), &req)
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(ctx, err, nil)
	}

	return utils.ReturnSuccessResponse(ctx, fiber.StatusOK, "Success create type", data)
}

// Update Type	godoc
// @Summary      	Update type
// @Description  	Update fasting type
// @Tags         	Admin
// @Accept       	json
// @Produce      	json
// @param 		 	payload body requests.TypeRequest true "JSON payload"
// @Param 			id path int true "Type ID"
// @Success      	200  {object}  utils.JSONResponse{data=transformer.TypeTransformer} "desc"
// @Failure      	400  {object}  utils.JSONResponse
// @Router       	/api/v1/admin/types/:id [put]
// @Security 		BasicAuth
func (c *AdminController) UpdateType(ctx *fiber.Ctx) error {
	var req requests.TypeRequest
	errorField, err := c.Core.Validator.Validate(ctx, &req)
	if err != nil {
		err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusUnprocessableEntity))
		return utils.ReturnErrorResponse(ctx, err, errorField)
	}

	idString := ctx.Params("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(ctx, err, nil)
	}

	data, err := c.AdminService.UpdateType(ctx.UserContext(), uint(id), &req)
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(ctx, err, nil)
	}

	return utils.ReturnSuccessResponse(ctx, fiber.StatusOK, "Success update type", data)
}

// Create Fasting	godoc
// @Summary      	Create fasting
// @Description  	Create fasting
// @Tags         	Admin
// @Accept       	json
// @Produce      	json
// @param 			payload body requests.FastingCreateUpdateRequest true "JSON payload"
// @Success      	200  {object}   utils.JSONResponse{data=transformer.FastingTransformer} "desc"
// @Failure      	400  {object}  utils.JSONResponse
// @Router       	/api/v1/admin/fastings [post]
// @Security 		BasicAuth
func (c *AdminController) CreateFasting(ctx *fiber.Ctx) error {
	var req requests.FastingCreateUpdateRequest
	errorField, err := c.Core.Validator.Validate(ctx, &req)
	if err != nil {
		err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusUnprocessableEntity))
		return utils.ReturnErrorResponse(ctx, err, errorField)
	}

	data, err := c.AdminService.CreateFasting(ctx.UserContext(), &req)
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(ctx, err, nil)
	}

	return utils.ReturnSuccessResponse(ctx, fiber.StatusOK, "Success create fasting", data)
}

// Update Fasting	godoc
// @Summary      	Update fasting
// @Description  	Update fasting
// @Tags         	Admin
// @Accept       	json
// @Produce      	json
// @param 		 	payload body requests.FastingCreateUpdateRequest true "JSON payload"
// @Param 			id path int true "Fasting ID"
// @Success      	200  {object}  utils.JSONResponse{data=transformer.FastingTransformer} "desc"
// @Failure      	400  {object}  utils.JSONResponse
// @Router       	/api/v1/admin/fastings/:id [put]
// @Security 		BasicAuth
func (c *AdminController) UpdateFasting(ctx *fiber.Ctx) error {
	var req requests.FastingCreateUpdateRequest
	errorField, err := c.Core.Validator.Validate(ctx, &req)
	if err != nil {
		err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusUnprocessableEntity))
		return utils.ReturnErrorResponse(ctx, err, errorField)
	}

	idString := ctx.Params("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(ctx, err, nil)
	}

	data, err := c.AdminService.UpdateFasting(ctx.UserContext(), &req, uint(id))
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(ctx, err, nil)
	}

	return utils.ReturnSuccessResponse(ctx, fiber.StatusOK, "Success update fasting", data)
}
