package db

import "github.com/jak103/powerplay/internal/models"

func (s session) GetKeyRecordById(id uint) (*models.KeyRecord, error) {
	var record *models.KeyRecord
	result := s.First(record, id)
	return resultOrError(record, result)
}
