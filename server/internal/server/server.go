package server

import (
	"github.com/gofiber/fiber/v2"
)

func Init() error {
	return nil
}

func Run() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8101")
}
