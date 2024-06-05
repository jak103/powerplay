package db

import (
	"github.com/jak103/powerplay/internal/models"
)

func (s Session) SaveGame(game models.Game) (*models.Game, error) {
	result := s.Connection.Create(&game)
	return resultOrError(&game, result)
}

func (s Session) SaveGames(games []models.Game) ([]models.Game, error) {
	result := s.Connection.CreateInBatches(games, len(games))
	return resultsOrError(games, result)
}

func (s Session) GetGame(id uint) (*models.Game, error) {
	game := models.Game{}
	result := s.Connection.First(&game, id)
	return resultOrError(&game, result)
}

func (s Session) GetGames(seasonId uint) (*[]models.Game, error) {
	games := make([]models.Game, 0)
	result := s.Connection.Where(&models.Game{SeasonID: seasonId}).Find(&games)
	return resultOrError(&games, result)
}

func (s Session) UpdateGame(game models.Game) (*models.Game, error) {
	result := s.Connection.Save(&game)
	return resultOrError(&game, result)
}

func (s Session) UpdateGames(games []models.Game) ([]models.Game, error) {
	result := s.Connection.Save(games)
	return resultsOrError(games, result)
}

func (s Session) DeleteGame(id uint) error {
	result := s.Connection.Delete(&models.Game{}, id)
	return result.Error
}

func (s Session) DeleteGames(seasonId uint) error {
	result := s.Connection.Where(&models.Game{SeasonID: seasonId}).Delete(&models.Game{})
	return result.Error
}
