package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
)

type session interface {
	GetUserByEmail(email string) (*models.User, error)
}

func newSession(c *fiber.Ctx) session {
	return db.GetSession(c)
}
