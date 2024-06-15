package db

import "github.com/jak103/powerplay/internal/models"

func (s session) GetSeasons() ([]models.Season, error) {
	seasons := make([]models.Season, 0)
	err := s.Preload("Leagues").Preload("Leagues.Teams").Find(&seasons)
	return resultsOrError(seasons, err)
}

func (s session) CreateSeason(season *models.Season) (*models.Season, error) {
	result := s.Create(season)
	return resultOrError(season, result)
}
