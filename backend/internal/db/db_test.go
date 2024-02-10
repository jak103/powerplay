package db

import (
	"testing"

	"github.com/jak103/powerplay/internal/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestResultOrErrorNominal(t *testing.T) {
	result := &gorm.DB{
		Error: nil,
	}

	var record *models.KeyRecord = &models.KeyRecord{
		UserId: 99,
	}

	r, err := resultOrError(record, result)

	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestResultOrErrorNoResult(t *testing.T) {
	result := &gorm.DB{
		Error: gorm.ErrRecordNotFound,
	}

	var record *models.KeyRecord = &models.KeyRecord{
		UserId: 99,
	}

	r, err := resultOrError(record, result)

	assert.Nil(t, r)
	assert.Nil(t, err)
}

func TestResultOrErrorError(t *testing.T) {
	result := &gorm.DB{
		Error: gorm.ErrDuplicatedKey,
	}

	var record *models.KeyRecord = &models.KeyRecord{
		UserId: 99,
	}

	r, err := resultOrError(record, result)

	assert.Nil(t, r)
	assert.NotNil(t, err)
}
