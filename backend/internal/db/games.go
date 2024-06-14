package db

import (
	"github.com/jak103/powerplay/internal/models"
)

func (s session) SaveGame(game models.Game) (*models.Game, error) {
	result := s.connection.Create(&game)
	return resultOrError(&game, result)
}

func (s session) SaveGames(games []models.Game) ([]models.Game, error) {
	result := s.connection.CreateInBatches(games, len(games))
	return resultsOrError(games, result)
}

func (s session) GetGame(id uint) (*models.Game, error) {
	game := models.Game{}
	result := s.connection.First(&game, id)
	return resultOrError(&game, result)
}

func (s session) GetGames(seasonId uint) (*[]models.Game, error) {
	games := make([]models.Game, 0)
	result := s.connection.Where(&models.Game{SeasonID: seasonId}).Find(&games)
	return resultOrError(&games, result)
}

func (s session) UpdateGame(id uint, game models.Game) (*models.Game, error) {
	// use the id to update the game
	game.ID = id
	result := s.connection.Save(&game)
	return resultOrError(&game, result)
}

func (s session) UpdateGames(games []models.Game) ([]models.Game, error) {
	result := s.connection.Save(games)
	return resultsOrError(games, result)
}

func (s session) DeleteGame(id uint) error {
	result := s.connection.Delete(&models.Game{}, id)
	return result.Error
}

func (s session) DeleteGames(seasonId uint) error {
	result := s.connection.Where(&models.Game{SeasonID: seasonId}).Delete(&models.Game{})
	return result.Error
}
