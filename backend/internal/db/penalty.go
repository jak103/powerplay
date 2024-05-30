package db

import "github.com/jak103/powerplay/internal/models"

func (s Session) GetPenalties() ([]models.Penalty, error) {
	penalties := make([]models.Penalty, 0)
	err := s.Connection.Preload("PenaltyType").Find(&penalties)
	return resultsOrError(penalties, err)
}

func (s Session) CreatePenalty(request *models.Penalty) error {
	result := s.Connection.Create(request)
	return result.Error
}

func (s Session) GetPenaltyTypes() ([]models.PenaltyType, error) {
	penaltyTypes := make([]models.PenaltyType, 0)
	err := s.Connection.Find(&penaltyTypes)
	return resultsOrError(penaltyTypes, err)
}
