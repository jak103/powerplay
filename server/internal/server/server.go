package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/leaguemanager/internal/config"
)

func Init(config *config.Config) error {
	return nil
}

func Run() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8101")
}
