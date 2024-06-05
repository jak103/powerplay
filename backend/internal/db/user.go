package db

import (
	"fmt"
	"slices"

	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/services/auth"
)

func (s Session) GetRefById(id uint) (*models.User, error) {
	var refUser *models.User
	result := s.Connection.First(&refUser, id)

	if result.Error != nil {
		return nil, fmt.Errorf("Error could not find the ref: %v in the db.", id)
	}

	if refUser != nil && !slices.Contains(refUser.Role, auth.Referee) {
		return nil, fmt.Errorf("The user %v is not a referee.", refUser.ID)
	}

	return resultOrError(refUser, result)
}
