package db

import (
	"github.com/jak103/powerplay/internal/models"
)

// GetGames retrieves all games from the database.
func (s session) GetGames() ([]models.League, error) {
	games := make([]models.League, 0)
	err := s.connection.Find(&games)
	return resultsOrError(games, err)
}

// GetGameByID retrieves a game by its ID from the database.
func (s *session) GetGameByID(gameID string) (*models.Game, error) {
	var game models.Game
	err := s.connection.First(&game, "id = ?", gameID).Error
	if err != nil {
		return nil, err
	}
	return &game, nil
}

// UpdateGame updates an existing game in the database.
func (s *session) UpdateGame(gameID string, game models.Game) error {
	var existingGame models.Game
	err := s.connection.First(&existingGame, "id = ?", gameID).Error
	if err != nil {
		return err
	}
	err = s.connection.Model(&existingGame).Updates(game).Error
	return err
}

// CreateGame creates a new game in the database.
func (s *session) CreateGame(game *models.Game) error {
	err := s.connection.Create(game).Error
	return err
}
