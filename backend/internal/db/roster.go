package db

import (
	"github.com/jak103/powerplay/internal/models"
)

func (s session) GetRosters() ([]models.Roster, error) {
	rosters := make([]models.Roster, 0)

	err := s.Preload("Players").Preload("Captain").Find(&rosters)

	return resultsOrError(rosters, err)
}

func (s session) CreateRoster(roster *models.Roster) (*models.Roster, error) {
	result := s.Create(roster)

	return resultOrError(roster, result)
}
