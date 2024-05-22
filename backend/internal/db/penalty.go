package db

import "github.com/jak103/powerplay/internal/models"

func (s session) GetPenalties() ([]models.Penalty, error) {
	penalties := make([]models.Penalty, 0)
	err := s.connection.Preload("PenaltyType").Find(&penalties)
	return resultsOrError(penalties, err)
}

func (s session) CreatePenalty(request *models.Penalty) error {
	result := s.connection.Create(request)
	return result.Error
}
