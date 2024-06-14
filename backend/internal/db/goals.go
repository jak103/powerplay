package db

import "github.com/jak103/powerplay/internal/models"

func (s session) SaveGoal(goal *models.Goal) (*models.Goal, error) {
	result := s.connection.Create(goal)
	return resultOrError(goal, result)
}

func (s session) GetGoals() ([]models.Goal, error) {
	goals := make([]models.Goal, 0)
	err := s.connection.Find(&goals)
	return resultsOrError(goals, err)
}
