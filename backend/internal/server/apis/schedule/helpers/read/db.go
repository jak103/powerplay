package read

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
)

// TODO not sure this is how you read from the db

func Games(c *fiber.Ctx, season string) ([]models.Game, error) {
	if c == nil || len(season) == 0 {
		return nil, errors.New("invalid uploads")
	}
	session := db.GetSession(c)
	var games []models.Game
	if err := session.Connection.Where("season = ?", season).Find(&games).Error; err != nil {
		return nil, err
	}
	return games, nil
}

func Leagues(c *fiber.Ctx, season string) ([]models.League, error) {
	if c == nil || len(season) == 0 {
		return nil, errors.New("invalid uploads")
	}
	session := db.GetSession(c)
	var leagues []models.League
	if err := session.Connection.Where("season = ?", season).Find(&leagues).Error; err != nil {
		return nil, err
	}
	return leagues, nil
}
