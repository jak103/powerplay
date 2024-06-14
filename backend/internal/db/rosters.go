package db

import (
	"github.com/jak103/powerplay/internal/models"
)

func (s Session) GetRosters() ([]models.Roster, error) {
	rosters := make([]models.Roster, 0)

	err := s.Connection.Preload("Players").Preload("Captain").Find(&rosters)

	return resultsOrError(rosters, err)
}

func (s Session) CreateRoster(roster *models.Roster) error {
	result := s.Connection.Create(roster)

	return result.Error
}

func (s Session) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}

	err := s.Connection.Find(&user, "email = ?", email)

	return resultOrError(user, err)
}

func (s Session) GetUserByEmails(emails []string) ([]*models.User, error) {
	users := make([]*models.User, 0)

	err := s.Connection.Where("email IN (?)", emails).Find(&users)

	return resultsOrError(users, err)
}
