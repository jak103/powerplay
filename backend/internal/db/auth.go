package db

import "github.com/jak103/powerplay/internal/models"

func (s *session) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	result := s.connection.Where("email = ?", email).First(user)

	return resultOrError(user, result)
}
