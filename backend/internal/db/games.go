package db

import (
	"github.com/jak103/powerplay/internal/models"
)

func (s session) SaveGame(game models.Game) (*models.Game, error) {
	result := s.Create(&game)
	return resultOrError(&game, result)
}

func (s session) SaveGames(games []models.Game) ([]models.Game, error) {
	result := s.CreateInBatches(games, len(games))
	return resultsOrError(games, result)
}

func (s session) GetGame(id uint) (*models.Game, error) {
	game := models.Game{}
	result := s.First(&game, id)
	return resultOrError(&game, result)
}

func (s session) GetGames(seasonId uint) (*[]models.Game, error) {
	games := make([]models.Game, 0)
	result := s.Where(&models.Game{SeasonID: seasonId}).Find(&games)
	return resultOrError(&games, result)
}

func (s session) GetGameById(id uint) (*models.Game, error) {
        game := &models.Game{}

        err := s.Find(&game, "id = ?", id)

        return resultOrError(game, err)
}

func (s session) UpdateGame(id uint, game models.Game) (*models.Game, error) {
	// use the id to update the game
	game.ID = id
	result := s.Save(&game)
	return resultOrError(&game, result)
}

func (s session) UpdateGames(games []models.Game) ([]models.Game, error) {
	result := s.Save(games)
	return resultsOrError(games, result)
}

func (s session) DeleteGame(id uint) error {
	result := s.Delete(&models.Game{}, id)
	return result.Error
}

func (s session) DeleteGames(seasonId uint) error {
	result := s.Where(&models.Game{SeasonID: seasonId}).Delete(&models.Game{})
	return result.Error
}
