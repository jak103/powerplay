package db

import (
	"github.com/jak103/powerplay/internal/models"
)

func (s session) GetUserByID(id uint) (*models.User, error) {
	user := &models.User{}

	err := s.Find(&user, "ID = ?", id)

	return resultOrError(user, err)
}

func (s session) GetUsersByIDs(ids []uint) ([]*models.User, error) {
	users := make([]*models.User, 0)

	err := s.Where("ID IN (?)", ids).Find(&users)

	return resultsOrError(users, err)
}
