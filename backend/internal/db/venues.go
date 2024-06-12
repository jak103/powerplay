package db

import "github.com/jak103/powerplay/internal/models"

func (s session) SaveVenue(venue *models.Venue) (*models.Venue, error) {
	result := s.connection.Create(venue)
	return resultOrError(venue, result)
}

func (s session) GetVenues() ([]models.Venue, error) {
	venues := make([]models.Venue, 0)
	err := s.connection.Find(&venues)
	return resultsOrError(venues, err)
}