package chat

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
)

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/hello", auth.Public, helloWorld)
}

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}
