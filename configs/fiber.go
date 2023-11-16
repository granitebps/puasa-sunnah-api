package config

import (
	"errors"
	"time"

	"github.com/ansel1/merry/v2"
	"github.com/goccy/go-json"
	"github.com/granitebps/puasa-sunnah-api/pkg/constants"
	"github.com/granitebps/puasa-sunnah-api/pkg/utils"
	"github.com/spf13/viper"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	return fiber.Config{
		AppName:     viper.GetString(constants.APP_NAME),
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		ReadTimeout: time.Second * 60,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError
			message := err.Error()

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
				message = e.Message
			}

			err = merry.Wrap(err, merry.WithHTTPCode(code), merry.WithUserMessage(message))

			return utils.ReturnErrorResponse(c, err, nil)
		},
	}
}
