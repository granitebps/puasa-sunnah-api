package controlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/services"
)

// ListFasting godoc
// @Summary      List Sunnah Fastings
// @Description  Get list of sunnah fasting
// @Tags         Fastings
// @Accept       json
// @Produce      json
// @Success      200  {object}   helpers.SuccessResponse{data=[]types.Fasting} "desc"
// @Failure      400  {object}  helpers.FailedResponse
// @Router       /api/v1/fastings [get]
func FastingsIndex(c *fiber.Ctx) error {
	data, err := services.FastingsGetAll()
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
