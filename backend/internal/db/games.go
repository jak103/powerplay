package db

import (
	"fmt"

	"github.com/jak103/powerplay/internal/models"
)

func (s Session) SaveGames(games []models.Game) ([]models.Game, error) {
	result := s.Connection.CreateInBatches(games, len(games))
	return resultsOrError(games, result)
}

func (s Session) GetGames(seasonId uint) (*[]models.Game, error) {
	games := make([]models.Game, 0)
	result := s.Connection.Where(&models.Game{SeasonID: seasonId}).Find(&games)
	return resultOrError(&games, result)
}

func (s Session) GetGame(gameId uint) (*models.Game, error) {
	var game *models.Game
	result := s.Connection.First(&game, gameId)

	if result.Error != nil {
		return nil, fmt.Errorf("Error could not find the ref: %v in the db.", id)
	}

	return resultOrError(game, result)
}

func (s Session) UpdateGames(games []models.Game) ([]models.Game, error) {
	result := s.Connection.Save(games)
	return resultsOrError(games, result)
}
