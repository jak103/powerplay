package db

import "github.com/jak103/powerplay/internal/models"

func (s session) SaveShotOnGoal(shotOnGoal *models.ShotOnGoal) (*models.ShotOnGoal, error) {
	result := s.connection.Create(shotOnGoal)
	return resultOrError(shotOnGoal, result)
}
