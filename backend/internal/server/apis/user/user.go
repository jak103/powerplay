package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
)

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/user", auth.Authenticated, getCurrentUser)
	apis.RegisterHandler(fiber.MethodPost, "/user", auth.Public, createUserAccount)

}

func getCurrentUser(c *fiber.Ctx) error {
	return nil
}

func createUserAccount(c *fiber.Ctx) error {
	return nil
}
