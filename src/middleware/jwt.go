package middleware

import (
	"github.com/ansel1/merry/v2"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/granitebps/puasa-sunnah-api/pkg/constants"
	"github.com/granitebps/puasa-sunnah-api/pkg/utils"
	"github.com/spf13/viper"
)

func Private() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(viper.GetString(constants.JWT_SECRET))},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			err = merry.Wrap(err, merry.WithHTTPCode(fiber.StatusUnauthorized), merry.WithUserMessage("Unauthenticated"))
			return utils.ReturnErrorResponse(c, err, nil)
		},
	})
}
