package db

import "github.com/jak103/powerplay/internal/models"

func (s *session) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	result := s.connection.Where("email = ?", email).First(user)

	return resultOrError(user, result)
}

func (s *session) CreateUser(user *models.User) (*models.User, error) {
	result := s.connection.Create(user)

	return resultOrError(user, result)
}
