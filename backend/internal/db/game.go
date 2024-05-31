package db

import (
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/utils/log"
)

// GetGames retrieves all games from the database.
func (s *session) GetGames() ([]models.Game, error) {
	games := make([]models.Game, 0)
	err := s.connection.Find(&games).Error
	if err != nil {
		log.WithErr(err).Alert("Failed to retrieve games from database")
		return nil, err
	}
	return games, nil
}

// GetGameByID retrieves a game by its ID from the database.
func (s *session) GetGameByID(gameID string) (*models.Game, error) {
	var game models.Game
	err := s.connection.First(&game, "id = ?", gameID).Error
	if err != nil {
		log.WithErr(err).Alert("Failed to retrieve game from database")
		return nil, err
	}
	return &game, nil
}

// UpdateGame updates an existing game in the database.
func (s *session) UpdateGame(gameID string, game models.Game) error {
	var existingGame models.Game
	err := s.connection.First(&existingGame, "id = ?", gameID).Error
	if err != nil {
		log.WithErr(err).Alert("Failed to find existing game in database")
		return err
	}
	err = s.connection.Model(&existingGame).Updates(game).Error
	if err != nil {
		log.WithErr(err).Alert("Failed to update game in database")
	}
	return err
}

// CreateGame creates a new game in the database.
func (s *session) CreateGame(game *models.Game) error {
	err := s.connection.Create(game).Error
	if err != nil {
		log.WithErr(err).Alert("Failed to create game in database")
	}
	return err
}
