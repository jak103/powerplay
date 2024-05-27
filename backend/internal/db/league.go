package db

import "github.com/jak103/powerplay/internal/models"

// GetLeagues todo: investigate: Struct session has methods on both value and pointer receivers. Such usage is not recommended by the Go Documentation.
func (s session) GetLeagues() ([]models.League, error) {
	leagues := make([]models.League, 0)
	err := s.connection.Find(&leagues)
	return resultsOrError(leagues, err)
}
