package db

import (
	"github.com/jak103/powerplay/internal/models"
)

func (s Session) GetTeams() ([]models.Team, error) {
	team := make([]models.Team, 0)
	err := s.Connection.Find(&team)
	return resultsOrError(team, err)
}

// GetTeamByID retrieves a team by its ID from the database.
func (s *Session) GetTeamByID(teamID string) (*models.Team, error) {
	var team models.Team
	err := s.Connection.First(&team, "id = ?", teamID).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

// UpdateTeam updates an existing team in the database.
func (s *Session) UpdateTeam(teamID string, team models.Team) error {
	var existingTeam models.Team
	err := s.Connection.First(&existingTeam, "id = ?", teamID).Error
	if err != nil {
		return err
	}
	err = s.Connection.Model(&existingTeam).Updates(team).Error
	return err
}

// CreateTeam creates a new team in the database.
func (s *Session) CreateTeam(team *models.Team) error {
	err := s.Connection.Create(team).Error
	return err
}
