package db

import "github.com/jak103/powerplay/internal/models"

func (s Session) GetKeyRecordById(id uint) (*models.KeyRecord, error) {
	var record *models.KeyRecord
	result := s.connection.First(record, id)
	return resultOrError(record, result)
}
