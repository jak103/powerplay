package hello

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/auth"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/utils/log"
)

func Init() {
	log.Debug("Registering hello handler")
	apis.RegisterHandler(fiber.MethodGet, "/hello", auth.NONE, handleGetHello)
}

func handleGetHello(c *fiber.Ctx) error {

	return c.SendString("Hello World")
}
