package controller

import (
	"github.com/ansel1/merry"
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/pkg/utils"
	"github.com/granitebps/puasa-sunnah-api/src/service"
)

type TypesController struct {
	TypesService *service.TypesService
}

func newTypesController(typesService *service.TypesService) *TypesController {
	return &TypesController{
		TypesService: typesService,
	}
}

// ListTypes godoc
// @Summary      List Types
// @Description  Get list of types
// @Tags         Types
// @Accept       json
// @Produce      json
// @Success      200  {object}   utils.JSONResponse{data=[]transformer.TypeTransformer} "desc"
// @Failure      400  {object}  utils.JSONResponse
// @Router       /api/v1/types [get]
func (c *TypesController) Index(ctx *fiber.Ctx) error {
	data, err := c.TypesService.GetAll(ctx.UserContext())
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(ctx, err, nil)
	}

	return utils.ReturnSuccessResponse(ctx, fiber.StatusOK, "Success", data)
}
