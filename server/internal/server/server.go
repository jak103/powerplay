package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/leaguemanager/internal/middleware"
)

func Init() error {
	return nil
}

func Run() {
	app := fiber.New()

	middleware.Setup(app)

	setupRoutes(app)

	app.Listen(":8101")
}
