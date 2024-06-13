package db

import (
	"github.com/jak103/powerplay/internal/models"
)

func (s session) GetTeams() ([]models.Team, error) {
	team := make([]models.Team, 0)
	err := s.Find(&team)
	return resultsOrError(team, err)
}

// GetTeamByID retrieves a team by its ID from the database.
func (s *session) GetTeamByID(teamID string) (*models.Team, error) {
	var team models.Team
	err := s.First(&team, "id = ?", teamID).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

// UpdateTeam updates an existing team in the database.
func (s *session) UpdateTeam(teamID string, team models.Team) error {
	var existingTeam models.Team
	err := s.First(&existingTeam, "id = ?", teamID).Error
	if err != nil {
		return err
	}
	err = s.Model(&existingTeam).Updates(team).Error
	return err
}

// CreateTeam creates a new team in the database.
func (s *session) CreateTeam(team *models.Team) (*models.Team, error) {
	result := s.Create(team)
	return resultOrError(team, result)
}
