package db

import "github.com/jak103/powerplay/internal/models"

// GetLeagues todo: investigate: Struct session has methods on both value and pointer receivers. Such usage is not recommended by the Go Documentation.
func (s session) GetLeagues(sortField, sortOrder string) ([]models.LeagueRecord, error) {
	if sortField == "" {
		sortField = "ID"
	}
	if sortOrder == "" {
		sortOrder = "ASC"
	}

	leagues := make([]models.LeagueRecord, 0)
	err := s.connection.Order(sortField + " " + sortOrder).Find(&leagues)
	return resultsOrError(leagues, err)
}

func (s session) GetLeaguesPaginated(offset, limit int, sortField, sortOrder string) ([]models.LeagueRecord, error) {
	if sortField == "" {
		sortField = "ID"
	}
	if sortOrder == "" {
		sortOrder = "asc"
	}
	leagues := make([]models.LeagueRecord, 0)
	err := s.connection.Preload("Teams").Offset(offset).Limit(limit).Order(sortField + " " + sortOrder).Find(&leagues)
	return resultsOrError(leagues, err)
}

func (s session) CreateLeague(request *models.LeagueRecord) error {
	result := s.connection.Create(request)
	return result.Error
}
