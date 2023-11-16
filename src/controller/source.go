package controller

import (
	"github.com/ansel1/merry"
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/pkg/utils"
	"github.com/granitebps/puasa-sunnah-api/src/service"
)

type SourceController struct {
	SourceService *service.SourceService
}

func newSourceController(sourceService *service.SourceService) *SourceController {
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
// @Success      200  {object}   utils.JSONResponse{data=[]types.Source} "desc"
// @Failure      400  {object}  utils.JSONResponse
// @Router       /api/v1/sources [get]
func (c *SourceController) Index(ctx *fiber.Ctx) error {
	data, err := c.SourceService.GetAll()
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(ctx, err, nil)
	}

	return utils.ReturnSuccessResponse(ctx, fiber.StatusOK, "Success", data)
}
