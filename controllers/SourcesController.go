package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/errors"
	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/services"
)

type SourceController struct {
	SourceService *services.SourceService
}

func newSourceController(sourceService *services.SourceService) *SourceController {
	return &SourceController{
		SourceService: sourceService,
	}
}

// ListSource godoc
// @Summary      List Sources
// @Description  Get list of sources
// @Tags         Sources
// @Accept       json
// @Produce      json
// @Success      200  {object}   helpers.SuccessResponse{data=[]types.Source} "desc"
// @Failure      400  {object}  helpers.FailedResponse
// @Router       /api/v1/sources [get]
func (c *SourceController) Index(ctx *fiber.Ctx) error {
	data, err := c.SourceService.GetAll()
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
