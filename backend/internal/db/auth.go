package db

import "github.com/jak103/powerplay/internal/models"

func (s *Session) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}

	result := s.connection.Where("username = ?", username).First(user)

	return resultOrError(user, result)
}
