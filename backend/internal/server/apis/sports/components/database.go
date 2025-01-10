package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
)

//go:generate moq -out database_mock_test.go . session
type session interface {
	CreateSeason(goal *models.Season) (*models.Season, error)
	GetSeasons() ([]models.Season, error)
}

type connection struct {
	newSession func(c *fiber.Ctx) session
}

var database = &connection{
	newSession: func(c *fiber.Ctx) session {
		return db.GetSession(c)
	},
}
