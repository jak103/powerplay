package db

import (
	"github.com/jak103/powerplay/internal/models"
)

func (s session) SaveVenue(venue *models.Venue) (*models.Venue, error) {
	result := s.Create(venue)
	return resultOrError(venue, result)
}

func (s session) GetVenues() ([]models.Venue, error) {
	venues := make([]models.Venue, 0)
	err := s.Find(&venues)
	return resultsOrError(venues, err)
}

func (s session) GetVenueById(id uint) (*models.Venue, error) {
	var venue models.Venue
	err := s.First(&venue, "id = ?", id)
	return resultOrError(&venue, err)
}
