package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/middleware"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/apis/hello"
	"github.com/jak103/powerplay/internal/utils/locals"
)

func Init() error {
	hello.Init()
	return nil
}

func Run() {
	app := fiber.New(fiber.Config{
		ErrorHandler: globalErrorHandler,
	})

	middleware.Setup(app)

	apis.SetupRoutes(app)

	app.Get("/test", func(c *fiber.Ctx) error { return c.SendString("test") })

	app.Listen(":9001")

	// Figure out websockets for chat and live scoring
	// TODO https://github.com/gofiber/contrib/tree/main/websocket
}

func globalErrorHandler(c *fiber.Ctx, err error) error {
	log := locals.Logger(c)

	log.WithErr(err).Error("Caught by global error handler")

	return c.Status(fiber.StatusNotFound).SendString("Not found")
}
