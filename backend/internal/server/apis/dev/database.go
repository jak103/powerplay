package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
)

//go:generate moq -out database_mock_test.go . session
type session interface {
	ResetDatabase() error
}

var newSession func(c *fiber.Ctx) session = func(c *fiber.Ctx) session {
	return db.GetSession(c)
}
