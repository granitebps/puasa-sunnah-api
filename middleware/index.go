package middleware

import (
	"time"

	"github.com/gofiber/contrib/fibersentry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/helmet/v2"
)

func InitMiddleware(app *fiber.App) *fiber.App {
	// Logger
	app.Use(logger.New())

	// Request ID
	app.Use(requestid.New())

	// ETag
	app.Use(etag.New())

	// Cache
	app.Use(cache.New(cache.Config{
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.OriginalURL()
		},
		Expiration: 1 * time.Hour,
	}))

	// Rate limit
	app.Use(limiter.New(limiter.Config{
		Max: 60,
	}))

	// Recover
	app.Use(recover.New())

	// Helmet
	app.Use(helmet.New())

	app.Use(fibersentry.New(fibersentry.Config{
		Repanic: true,
	}))

	return app
}
