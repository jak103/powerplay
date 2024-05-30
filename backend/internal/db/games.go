package db

import (
	"github.com/jak103/powerplay/internal/models"
)

func (s Session) SaveGames(games []models.Game) ([]models.Game, error) {
	// TODO make this happy
	result := s.Connection.CreateInBatches(games, len(games))

	return resultsOrError(games, result)
}

func (s Session) GetGamesBySeason(seasonId uint) (*[]models.Game, error) {
        games := make([]models.Game, 0)
        result := s.Connection.Where(&models.Game{SeasonID: seasonId}).Find(&games)
        return resultOrError(&games, result)
}

// TODO finish this
