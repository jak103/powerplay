package db

import "github.com/jak103/powerplay/internal/models"

func (s Session) SaveGoal(goal *models.Goal) (*models.Goal, error) {
	result := s.Connection.Create(goal)
	return resultOrError(goal, result)
}

func (s Session) GetGoals() ([]models.Goal, error) {
	goals := make([]models.Goal, 0)
	err := s.Connection.Find(&goals)
	return resultsOrError(goals, err)
}
