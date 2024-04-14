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

	var user *models.User = &models.User{
		DbModel: models.DbModel{
			ID: 99,
		},
	}

	r, err := resultOrError(user, result)

	assert.NotNil(t, r)
	assert.Nil(t, err)
}

func TestResultOrErrorNoResult(t *testing.T) {
	result := &gorm.DB{
		Error: gorm.ErrRecordNotFound,
	}

	var user *models.User = &models.User{
		DbModel: models.DbModel{
			ID: 99,
		},
	}

	r, err := resultOrError(user, result)

	assert.Nil(t, r)
	assert.Nil(t, err)
}

func TestResultOrErrorError(t *testing.T) {
	result := &gorm.DB{
		Error: gorm.ErrDuplicatedKey,
	}

	var user *models.User = &models.User{
		DbModel: models.DbModel{
			ID: 99,
		},
	}

	r, err := resultOrError(user, result)

	assert.Nil(t, r)
	assert.NotNil(t, err)
}
