package write

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
)

func ToDb(c *fiber.Ctx, games []models.Game) error {
	if c == nil || len(games) == 0 {
		return errors.New("invalid input")
	}
	session := db.GetSession(c)

	for _, game := range games {
		if err := session.Connection.Save(&game).Error; err != nil {
			return err
		}
	}

	return nil
}
