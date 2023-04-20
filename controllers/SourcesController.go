package controlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/errors"
	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/services"
)

// ListSource godoc
// @Summary      List Sources
// @Description  Get list of sources
// @Tags         Sources
// @Accept       json
// @Produce      json
// @Success      200  {object}   helpers.SuccessResponse{data=[]types.Source} "desc"
// @Failure      400  {object}  helpers.FailedResponse
// @Router       /api/v1/sources [get]
func SourcesIndex(c *fiber.Ctx) error {
	data, err := services.SourcesGetAll()
	if err != nil {
		return helpers.FailedAPIResponse(
			c,
			errors.WrapUserMessageAndCode(err, "", 0),
			nil,
		)
	}

	return helpers.SuccessAPIResponse(
		c,
		"Success",
		fiber.StatusOK,
		data,
	)
}
