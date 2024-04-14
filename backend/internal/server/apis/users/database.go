package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
)

//go:generate moq -out database_mock_test.go . session
type session interface {
	GetUserByEmail(email string) (*models.User, error)
	CreateUser(email, passwordHash, salt string) error
	UpdateUser(user *models.User) (*models.User, error)
	GetUserById(id int) (*models.User, error)
}

var newSession func(c *fiber.Ctx) session = realSession

func realSession(c *fiber.Ctx) session {
	return db.GetSession(c)
}
