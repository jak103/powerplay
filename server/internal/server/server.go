package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/middleware"
)

func Init() error {
	return nil
}

func Run() {
	app := fiber.New()

	middleware.Setup(app)

	setupRoutes(app)

	app.Listen(":9001")

	// Figure out websockets for chat and live scoring
	// TODO https://github.com/gofiber/contrib/tree/main/websocket
}
