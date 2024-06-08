package db

import "github.com/jak103/powerplay/internal/models"

func (s Session) GetSeasons() ([]models.Season, error) {
	seasons := make([]models.Season, 0)
	err := s.Connection.Preload("Leagues").Find(&seasons)
	return resultsOrError(seasons, err)
}

func (s Session) SaveSeason(season *models.Season) (*models.Season, error) {
	result := s.Connection.Create(season)
	return resultOrError(season, result)
}
