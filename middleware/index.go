package middleware

import (
	"log"
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/gofiber/contrib/fibersentry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func InitMiddleware(app *fiber.App) *fiber.App {
	// Logger
	// now := time.Now().Format("2006-02-01")
	// logFileName := "./logs/" + now + ".log"
	// file, err := os.OpenFile(logFileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }
	// defer file.Close()
	// loggerConfig := logger.Config{
	// 	Output: file,
	// }
	// app.Use(logger.New(loggerConfig))
	app.Use(logger.New())

	// Request ID
	app.Use(requestid.New())

	// ETag
	app.Use(etag.New())

	// Cache
	app.Use(cache.New())

	// Rate limit
	app.Use(limiter.New(limiter.Config{
		Max: 60,
	}))

	// Recover
	app.Use(recover.New())

	// Sentry
	err := sentry.Init(sentry.ClientOptions{
		Dsn:   os.Getenv("SENTRY_DSN"),
		Debug: true,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	app.Use(fibersentry.New(fibersentry.Config{
		Repanic: true,
	}))

	return app
}
