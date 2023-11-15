package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/errors"
	"github.com/granitebps/puasa-sunnah-api/helpers"
	"github.com/granitebps/puasa-sunnah-api/services"
)

// ListTypes godoc
// @Summary      List Types
// @Description  Get list of types
// @Tags         Types
// @Accept       json
// @Produce      json
// @Success      200  {object}   helpers.SuccessResponse{data=[]types.Type} "desc"
// @Failure      400  {object}  helpers.FailedResponse
// @Router       /api/v1/types [get]
func TypesIndex(c *fiber.Ctx) error {
	data, err := services.TypesGetAll()
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
