package db

import "github.com/jak103/powerplay/internal/models"

func (s session) SaveGoal(goal *models.Goal) (*models.Goal, error) {
	result := s.connection.Create(goal)
	return resultOrError(goal, result)
}
