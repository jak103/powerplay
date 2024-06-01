package db

import (
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

func (s Session) UpdateGames(games []models.Game) ([]models.Game, error) {
	result := s.Connection.Save(games)
	return resultsOrError(games, result)
}
