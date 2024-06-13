package db

import "github.com/jak103/powerplay/internal/models"

func (s session) GetPenalties() ([]models.Penalty, error) {
	penalties := make([]models.Penalty, 0)
	err := s.Preload("PenaltyType").Find(&penalties)
	return resultsOrError(penalties, err)
}

func (s session) CreatePenalty(request *models.Penalty) error {
	result := s.Create(request)
	return result.Error
}

func (s session) GetPenaltyTypes() ([]models.PenaltyType, error) {
	penaltyTypes := make([]models.PenaltyType, 0)
	err := s.Find(&penaltyTypes)
	return resultsOrError(penaltyTypes, err)
}
