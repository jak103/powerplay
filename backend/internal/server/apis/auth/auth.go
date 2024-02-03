package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/auth"
	"github.com/jak103/powerplay/internal/server/apis"
)

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/auth", auth.NONE, postAuthHandler)

	apis.RegisterHandler(fiber.MethodGet, "/user", auth.JWT, getCurrentUser)
	apis.RegisterHandler(fiber.MethodPost, "/user", auth.NONE, createUserAccount)
}

func postAuthHandler(c *fiber.Ctx) error {
	return nil
}

func getCurrentUser(c *fiber.Ctx) error {
	return nil
}

func createUserAccount(c *fiber.Ctx) error {
	return nil
}
