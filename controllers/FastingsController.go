package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/errors"
	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/requests"
	"github.com/granitebps/puasa-sunnah-api/services"
)

type FastingController struct {
	FastingService *services.FastingService
}

func newFastingController(fastingService *services.FastingService) *FastingController {
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
// @Success      200  {object}   helpers.SuccessResponse{data=[]types.Fasting} "desc"
// @Failure      400  {object}  helpers.FailedResponse
// @Router       /api/v1/fastings [get]
func (c *FastingController) Index(ctx *fiber.Ctx) error {
	f := requests.FastingRequest{}

	if err := ctx.QueryParser(&f); err != nil {
		return err
	}

	data, err := c.FastingService.GetAll(f)
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
