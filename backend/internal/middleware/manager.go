package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/fiber/v2/utils"

	"github.com/jak103/powerplay/internal/config"
	"github.com/jak103/powerplay/internal/middleware/logger"
	"github.com/jak103/powerplay/internal/middleware/metrics"
	appUtils "github.com/jak103/powerplay/internal/utils"
	"github.com/jak103/powerplay/internal/utils/constants"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func Setup(app *fiber.App) {
	app.Use(metrics.New())

	app.Use(helmet.New())

	// Request ID https://docs.gofiber.io/api/middleware/requestid
	app.Use(requestid.New(requestid.Config{
		Generator:  utils.UUIDv4,
		ContextKey: constants.RequestIdLocal,
	}))

	app.Use(logger.New())

	app.Use(recover.New(recover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: appUtils.StackTraceHandler,
	}))

	// CSRF https://docs.gofiber.io/api/middleware/csrf
	if config.Vars.Env != constants.Local && config.Vars.Env != constants.Test { // Only turn on CSRF in deployed environments
		app.Use(csrf.New(csrf.Config{
			CookieHTTPOnly: true,
			CookieSecure:   true,
			KeyLookup:      "cookie:csrf_",
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				locals.Logger(c).WithErr(err).Info("CRSF middlware failed with error")
				return responder.BadRequest(c, "Could not find CRSF token, refresh page")
			},
		}))
	}

	// Compression https://docs.gofiber.io/api/middleware/compress
	app.Use(compress.New(compress.Config{
		Level: compress.LevelDefault,
	}))

	// TODO Setup CORS
	// if config.Devsite.Env == constants.Local || config.Devsite.Env == constants.Test {
	// 	log.Error("Settting permissive CORS")
	// 	// CORS https://docs.gofiber.io/api/middleware/cors
	// 	app.Use(cors.New(cors.Config{
	// 		AllowOrigins:     "http://localhost:5173",
	// 		AllowCredentials: true,
	// 		AllowMethods:     "POST, GET, OPTIONS, PUT, DELETE",
	// 		AllowHeaders:     "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Cookie",
	// 		ExposeHeaders:    "Set-Cookie",
	// 	}))
	// }

	// TODO rate limiter https://github.com/gofiber/fiber/tree/v2/middleware/limiter
	// TODO otel traces https://github.com/gofiber/contrib/tree/main/otelfiber
}
