package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
)

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/auth", auth.Public, postAuthHandler)
}

func postAuthHandler(c *fiber.Ctx) error {
	return nil
}
