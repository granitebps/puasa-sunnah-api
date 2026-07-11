package middleware

import (
	"time"

	"github.com/ansel1/merry/v2"
	_ "github.com/granitebps/puasa-sunnah-api/docs"
	"github.com/granitebps/puasa-sunnah-api/pkg/constants"
	"github.com/granitebps/puasa-sunnah-api/pkg/core"
	utilsPkg "github.com/granitebps/puasa-sunnah-api/pkg/utils"

	fibernewrelic "github.com/gofiber/contrib/v3/newrelic"
	fibersentry "github.com/gofiber/contrib/v3/sentry"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cache"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/etag"
	"github.com/gofiber/fiber/v3/middleware/helmet"
	"github.com/gofiber/fiber/v3/middleware/limiter"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"
	"github.com/gofiber/fiber/v3/middleware/timeout"
	fiberutils "github.com/gofiber/utils/v2"

	swaggo "github.com/gofiber/contrib/v3/swaggo"
)

const (
	apiClientHeader          = "X-API-Client"
	officialAPIClient        = "official-app"
	requestSourceAttribute   = "request_source"
	requestSourceOfficialApp = "official_app"
	requestSourcePublic      = "public"
)

func classifyRequestSource(apiClient string) string {
	if apiClient == officialAPIClient {
		return requestSourceOfficialApp
	}

	return requestSourcePublic
}

func addRequestSourceAttribute(c fiber.Ctx) error {
	if txn := fibernewrelic.FromContext(c); txn != nil {
		txn.AddAttribute(requestSourceAttribute, classifyRequestSource(c.Get(apiClientHeader)))
	}

	return c.Next()
}

func SetupMiddleware(a *fiber.App, c *core.Core) {
	// Fiber Middleware
	a.Use(logger.New())
	a.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	// Swagger
	a.Get("/swagger/*", swaggo.HandlerDefault)

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
		LimitReached: func(c fiber.Ctx) error {
			err := merry.New("Too many requests.", merry.WithHTTPCode(fiber.StatusTooManyRequests), merry.WithUserMessage("Too many requests."))
			return utilsPkg.ReturnErrorResponse(c, err, nil)
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
	a.Use(addRequestSourceAttribute)

	a.Use(timeout.New(func(c fiber.Ctx) error {
		return c.Next()
	}, timeout.Config{Timeout: constants.TIMEOUT}))

	// Response Cache
	a.Use(cache.New(cache.Config{
		Expiration: 24 * time.Hour,
		Storage:    c.Cache.RedisStorage,
		KeyGenerator: func(c fiber.Ctx) string {
			return fiberutils.CopyString(c.OriginalURL())
		},
	}))
}
