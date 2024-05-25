package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/middleware"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/utils/locals"

	// Blank imports for apis to cause init functions to run
	_ "github.com/jak103/powerplay/internal/server/apis/auth"
	_ "github.com/jak103/powerplay/internal/server/apis/chat"
	_ "github.com/jak103/powerplay/internal/server/apis/game"
	_ "github.com/jak103/powerplay/internal/server/apis/notifications"
	_ "github.com/jak103/powerplay/internal/server/apis/stats"
	_ "github.com/jak103/powerplay/internal/server/apis/user"
)

func Run() {
	app := fiber.New(fiber.Config{
		ErrorHandler:          globalErrorHandler,
		DisableStartupMessage: true,
	})

	middleware.Setup(app)

	apis.SetupRoutes(app)

	app.Static("/", "/app/static") // TODO make this an env var?

	app.Listen(":9001")
}

func globalErrorHandler(c *fiber.Ctx, err error) error {
	log := locals.Logger(c)

	log.WithErr(err).Error("Caught by global error handler")

	return c.Status(fiber.StatusNotFound).SendString("Not found")
}
