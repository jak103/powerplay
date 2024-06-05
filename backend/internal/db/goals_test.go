package db

import (
	"github.com/jak103/powerplay/internal/models"
)

func (s *dbTestingSuite) TestSaveGoal() {
	goal := models.Goal{
		GameId: 1,
		UserId: 1,
		Period: 2,
	}

	g, err := s.session.SaveGoal(&goal)
	s.Nil(err)

	s.Equal(uint(1), g.ID)

	goals, err := s.session.GetGoals()
	s.Nil(err)

	s.Len(goals, 1)
}
