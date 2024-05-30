package db

import "github.com/jak103/powerplay/internal/models"

func (s Session) SaveShotOnGoal(shotOnGoal *models.ShotOnGoal) (*models.ShotOnGoal, error) {
	result := s.Connection.Create(shotOnGoal)
	return resultOrError(shotOnGoal, result)
}
