package controller

import (
	"github.com/ansel1/merry"
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/pkg/utils"
	"github.com/granitebps/puasa-sunnah-api/requests"
	"github.com/granitebps/puasa-sunnah-api/src/service"
)

type FastingController struct {
	FastingService *service.FastingService
}

func newFastingController(fastingService *service.FastingService) *FastingController {
	return &FastingController{
		FastingService: fastingService,
	}
}

// ListFasting godoc
// @Summary      List Sunnah Fastings
// @Description  Get list of sunnah fasting
// @Tags         Fastings
// @Accept       json
// @Produce      json
// @Param 		 type_id query int false "Type ID"
// @Param 		 category_id query int false "Category ID"
// @Param 		 day query int false "Day in month"
// @Param 		 month query int false "Month"
// @Param 		 Year query int false "Year"
// @Success      200  {object}   utils.JSONResponse{data=[]types.Fasting} "desc"
// @Failure      400  {object}  utils.JSONResponse
// @Router       /api/v1/fastings [get]
func (c *FastingController) Index(ctx *fiber.Ctx) error {
	f := requests.FastingRequest{}

	if err := ctx.QueryParser(&f); err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(ctx, err, nil)
	}

	data, err := c.FastingService.GetAll(f)
	if err != nil {
		err = merry.Wrap(err)
		return utils.ReturnErrorResponse(ctx, err, nil)
	}

	return utils.ReturnSuccessResponse(ctx, fiber.StatusOK, "Success", data)
}
