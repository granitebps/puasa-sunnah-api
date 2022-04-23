package controlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/requests"
	"github.com/granitebps/puasa-sunnah-api/services"
)

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
func FastingsIndex(c *fiber.Ctx) error {
	f := requests.FastingRequest{}

	if err := c.QueryParser(&f); err != nil {
		return err
	}

	data, err := services.FastingsGetAll(f)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(helpers.FailedAPIResponse(
			err.Error(),
			http.StatusBadRequest,
		))
	}

	return c.Status(http.StatusOK).JSON(helpers.SuccessAPIResponse(
		"Success",
		http.StatusOK,
		data,
	))
}
