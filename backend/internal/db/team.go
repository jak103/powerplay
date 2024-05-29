package db

import (
	"github.com/jak103/powerplay/internal/models"
)

func (s session) GetTeams() ([]models.League, error) {
	team := make([]models.League, 0)
	err := s.connection.Find(&team)
	return resultsOrError(team, err)
}

// GetTeamByID retrieves a team by its ID from the database.
func (s *session) GetTeamByID(teamID string) (*models.Team, error) {
	var team models.Team
	err := s.connection.First(&team, "id = ?", teamID).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

// UpdateTeam updates an existing team in the database.
func (s *session) UpdateTeam(teamID string, team models.Team) error {
	var existingTeam models.Team
	err := s.connection.First(&existingTeam, "id = ?", teamID).Error
	if err != nil {
		return err
	}
	err = s.connection.Model(&existingTeam).Updates(team).Error
	return err
}

// CreateTeam creates a new team in the database.
func (s *session) CreateTeam(team *models.Team) error {
	err := s.connection.Create(team).Error
	return err
}

// GetTeam todo: investigate: Struct session has methods on both value and pointer receivers. Such usage is not recommended by the Go Documentation.
func (s session) GetTeam() ([]models.Team, error) {
	Team := make([]models.Team, 0)
	err := s.connection.Find(&Team)
	return resultsOrError(Team, err)
}
