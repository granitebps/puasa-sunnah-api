package middleware

import (
	"time"

	"github.com/ansel1/merry/v2"
	_ "github.com/granitebps/puasa-sunnah-api/docs"
	"github.com/granitebps/puasa-sunnah-api/pkg/constants"
	"github.com/granitebps/puasa-sunnah-api/pkg/core"
	"github.com/granitebps/puasa-sunnah-api/pkg/utils"

	"github.com/gofiber/contrib/fibernewrelic"
	"github.com/gofiber/contrib/fibersentry"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/middleware/timeout"
	"github.com/gofiber/swagger"
)

func SetupMiddleware(a *fiber.App, c *core.Core) {
	// Fiber Middleware
	a.Use(logger.New())
	a.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	// Swagger
	// We need to put swagger middleware in here to prevent collision with security middleware
	a.Get("/swagger/*", swagger.HandlerDefault)

	// Another Fiber Middleware
	a.Use(etag.New())
	a.Use(compress.New())
	a.Use(cors.New())
	a.Use(requestid.New())
	a.Use(helmet.New())
	a.Use(limiter.New(limiter.Config{
		Max:               100,
		Expiration:        1 * time.Second,
		LimiterMiddleware: limiter.SlidingWindow{},
		LimitReached: func(c *fiber.Ctx) error {
			err := merry.New("Too many requests.", merry.WithHTTPCode(fiber.StatusTooManyRequests), merry.WithUserMessage("Too many requests."))
			return utils.ReturnErrorResponse(c, err, nil)
		},
	}))

	// Sentry
	a.Use(fibersentry.New(fibersentry.Config{
		Repanic: true,
	}))

	// Newrelic
	a.Use(fibernewrelic.New(fibernewrelic.Config{
		Application: c.Newrelic,
	}))

	a.Use(
		timeout.NewWithContext(func(c *fiber.Ctx) error {
			return c.Next()
		}, constants.TIMEOUT),
	)

	// Uncomment these code if you want to implement https://docs.gofiber.io/api/middleware/monitor
	// a.Get("/metrics", monitor.New(monitor.Config{
	// 	Title: fmt.Sprintf("%s Monitor", c.AppName),
	// }))

	// Response Cache
	a.Use(cache.New(cache.Config{
		Expiration:   24 * time.Hour, // 24 hour
		Storage:      c.Cache.RedisStorage,
		CacheControl: true,
	}))
}
