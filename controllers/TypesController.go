package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/errors"
	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/services"
)

type TypesController struct {
	TypesService *services.TypesService
}

func newTypesController(typesService *services.TypesService) *TypesController {
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
// @Success      200  {object}   helpers.SuccessResponse{data=[]types.Type} "desc"
// @Failure      400  {object}  helpers.FailedResponse
// @Router       /api/v1/types [get]
func (c *TypesController) Index(ctx *fiber.Ctx) error {
	data, err := c.TypesService.GetAll()
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
