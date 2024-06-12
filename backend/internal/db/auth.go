package db

import "github.com/jak103/powerplay/internal/models"

func (s *session) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}

	result := s.Where("username = ?", username).First(user)

	return resultOrError(user, result)
}

func (s *session) CreateUser(user *models.User) (*models.User, error) {
	result := s.connection.Create(user)

	return resultOrError(user, result)
}
