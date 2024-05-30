package db

import (
	"github.com/jak103/powerplay/internal/models"
)

func (s Session) SaveGames(games []models.Game) ([]models.Game, error) {
	// TODO make this happy
	result := s.Connection.CreateInBatches(games, len(games))

	return resultsOrError(games, result)
}

// TODO finish this
