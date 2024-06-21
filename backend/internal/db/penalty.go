package db

import "github.com/jak103/powerplay/internal/models"

func (s session) GetPenalties() ([]models.Penalty, error) {
	penalties := make([]models.Penalty, 0)
	err := s.Preload("PenaltyType").Find(&penalties)
	return resultsOrError(penalties, err)
}

func (s session) GetPenaltyByID(id string) (*models.Penalty, error) {
	var penalty *models.Penalty
	err := s.Preload("PenaltyType").First(&penalty, "id = ?", id)
	return resultOrError(penalty, err)
}

func (s session) GetPenaltiesByGameID(gameID string) ([]models.Penalty, error) {
	penalties := make([]models.Penalty, 0)
	err := s.Preload("PenaltyType").Where("game_id = ?", gameID).Find(&penalties)
	return resultsOrError(penalties, err)
}

func (s session) GetPenaltiesByTeamID(teamID string) ([]models.Penalty, error) {
	penalties := make([]models.Penalty, 0)
	err := s.Preload("PenaltyType").Where("team_id = ?", teamID).Find(&penalties)
	return resultsOrError(penalties, err)
}

func (s session) GetPenaltiesByPlayerID(playerID string) ([]models.Penalty, error) {
	penalties := make([]models.Penalty, 0)
	err := s.Preload("PenaltyType").Where("player_id = ?", playerID).Find(&penalties)
	return resultsOrError(penalties, err)
}

func (s session) CreatePenalty(penalty *models.Penalty) (*models.Penalty, error) {
	err := s.Create(penalty)
	return resultOrError(penalty, err)
}

func (s session) UpdatePenalty(penalty *models.Penalty) (*models.Penalty, error) {
	err := s.Save(penalty)
	return resultOrError(penalty, err)
}

func (s session) DeletePenalty(penalty *models.Penalty) error {
	return s.Delete(penalty).Error
}
