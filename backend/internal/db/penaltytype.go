package db

import "github.com/jak103/powerplay/internal/models"

func (s session) GetPenaltyTypes() ([]models.PenaltyType, error) {
	penaltyTypes := make([]models.PenaltyType, 0)
	err := s.Find(&penaltyTypes)
	return resultsOrError(penaltyTypes, err)
}

func (s session) GetPenaltyTypeByID(id string) (*models.PenaltyType, error) {
	var penaltyType *models.PenaltyType
	err := s.First(&penaltyType, "id = ?", id)
	return resultOrError(penaltyType, err)
}

func (s session) CreatePenaltyType(penaltyType *models.PenaltyType) (*models.PenaltyType, error) {
	err := s.Create(penaltyType)
	return resultOrError(penaltyType, err)
}

func (s session) UpdatePenaltyType(penaltyType *models.PenaltyType) (*models.PenaltyType, error) {
	err := s.Save(penaltyType)
	return resultOrError(penaltyType, err)
}

func (s session) DeletePenaltyType(penaltyType *models.PenaltyType) error {
	return s.Delete(penaltyType).Error
}
