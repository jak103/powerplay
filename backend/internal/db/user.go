package db

import (
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/utils/log"
)

func (s session) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	result := s.connection.Where("email = ?", email).First(user)

	return resultOrError(user, result)
}

func (s session) GetUserById(id int) (*models.User, error) {
	user := &models.User{}

	result := s.connection.First(user, id)

	return resultOrError(user, result)
}

func (s session) CreateUser(email, passwordHash, salt string) error {
	newUser := models.User{
		Email:          email,
		HashedPassword: passwordHash,
		Verified:       false,
		Salt:           salt,
	}

	log.Debug("newUser.Email: %s", newUser.Email)

	result := s.connection.Create(&newUser)

	return result.Error
}

func (s session) UpdateUser(user *models.User) (*models.User, error) {
	result := s.connection.Updates(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return s.GetUserById(int(user.ID))
}
