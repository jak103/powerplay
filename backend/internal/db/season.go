package db

import "github.com/jak103/powerplay/internal/models"

func (s session) GetSeasons() ([]models.Season, error) {
	seasons := make([]models.Season, 0)
	err := s.connection.Preload("Leagues").Find(&seasons)
	return resultsOrError(seasons, err)
}

func (s session) SaveSeason(season *models.Season) (*models.Season, error) {
	result := s.connection.Create(season)
	return resultOrError(season, result)
}
