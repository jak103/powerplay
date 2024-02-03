package chat

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/auth"
	"github.com/jak103/powerplay/internal/server/apis"
)

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/hello", auth.NONE, helloWorld)
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello Worl")
}
