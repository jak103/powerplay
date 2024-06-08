package db

import "github.com/jak103/powerplay/internal/models"

// GetLeagues todo: investigate: Struct session has methods on both value and pointer receivers. Such usage is not recommended by the Go Documentation.
func (s Session) GetLeagues() ([]models.League, error) {
	leagues := make([]models.League, 0)
	err := s.Connection.Find(&leagues)
	return resultsOrError(leagues, err)
}

func (s Session) GetLeaguesBySeason(seasonId uint) ([]models.League, error) {
	leagues := make([]models.League, 0)
	err := s.Connection.Where(&models.League{SeasonID: seasonId}).Find(&leagues)
	return resultsOrError(leagues, err)
}

func (s Session) CreateLeague(request *models.League) error {
	result := s.Connection.Create(request)
	return result.Error
}
