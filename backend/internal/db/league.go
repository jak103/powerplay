package db

import "github.com/jak103/powerplay/internal/models"

// GetLeagues todo: investigate: Struct session has methods on both value and pointer receivers. Such usage is not recommended by the Go Documentation.
func (s *session) GetLeagues(sortField, sortOrder string) ([]models.League, error) {
	if sortField == "" {
		sortField = "ID"
	}
	if sortOrder == "" {
		sortOrder = "ASC"
	}

	leagues := make([]models.League, 0)
	err := s.Order(sortField + " " + sortOrder).Find(&leagues)
	return resultsOrError(leagues, err)
}

func (s session) GetLeaguesPaginated(offset, limit int, sortField, sortOrder string) ([]models.League, error) {
	if sortField == "" {
		sortField = "ID"
	}
	if sortOrder == "" {
		sortOrder = "asc"
	}
	leagues := make([]models.League, 0)
	err := s.Offset(offset).Limit(limit).Order(sortField + " " + sortOrder).Find(&leagues)
	return resultsOrError(leagues, err)
}

func (s session) CreateLeague(league *models.League) (*models.League, error) {
	result := s.Create(league)
	return resultOrError(league, result)
}

func (s session) GetLeaguesBySeason(seasonId int) ([]models.League, error) {
	leagues := make([]models.League, 0)
	result := s.Preload("Teams").Where("season_id = ?", seasonId).Find(&leagues)

	return resultsOrError(leagues, result)
}
