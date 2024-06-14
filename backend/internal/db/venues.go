package db

import "github.com/jak103/powerplay/internal/models"

func (s Session) SaveVenue(venue *models.Venue) (*models.Venue, error) {
	result := s.Connection.Create(venue)
	return resultOrError(venue, result)
}

func (s Session) GetVenues() ([]models.Venue, error) {
	venues := make([]models.Venue, 0)
	err := s.Connection.Find(&venues)
	return resultsOrError(venues, err)
}
