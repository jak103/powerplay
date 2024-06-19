package db

import "github.com/jak103/powerplay/internal/models"

func (s *dbTestingSuite) TestSaveGame() {
	game := models.Game{
		SeasonID:   1,
		HomeTeamID: 1,
		AwayTeamID: 1,
	}

	g, err := s.session.SaveGame(game)
	s.Nil(err)
	s.Equal(uint(1), g.ID)
	s.Equal(uint(1), g.SeasonID)
	s.Equal(uint(1), g.HomeTeamID)
	s.Equal(uint(1), g.AwayTeamID)

	games, err := s.session.GetGames(1)
	s.Nil(err)
	s.Len(games, 1)
}

func (s *dbTestingSuite) TestSaveGames() {
	games := []models.Game{
		{
			SeasonID:   1,
			HomeTeamID: 1,
			AwayTeamID: 1,
		},
		{
			SeasonID:   1,
			HomeTeamID: 2,
			AwayTeamID: 2,
		},
		{
			SeasonID:   1,
			HomeTeamID: 3,
			AwayTeamID: 3,
		},
	}

	g, err := s.session.SaveGames(games)
	s.Nil(err)
	s.Len(g, 3)

	other, err := s.session.GetGames(1)
	s.Nil(err)
	s.Len(other, 3)
}

func (s *dbTestingSuite) TestGetGame() {
	game := models.Game{
		SeasonID:   1,
		HomeTeamID: 1,
		AwayTeamID: 1,
	}

	g, err := s.session.SaveGame(game)
	s.Nil(err)
	s.NotNil(g)

	other, err := s.session.GetGame(1)
	s.Nil(err)
	s.Equal(uint(1), other.ID)
	s.Equal(uint(1), other.SeasonID)
	s.Equal(uint(1), other.HomeTeamID)
	s.Equal(uint(1), other.AwayTeamID)
}

func (s *dbTestingSuite) TestGetGames() {
	games := []models.Game{
		{
			SeasonID:   1,
			HomeTeamID: 1,
			AwayTeamID: 1,
		},
		{
			SeasonID:   1,
			HomeTeamID: 2,
			AwayTeamID: 2,
		},
		{
			SeasonID:   1,
			HomeTeamID: 3,
			AwayTeamID: 3,
		},
	}

	g, err := s.session.SaveGames(games)
	s.Nil(err)
	s.Len(g, 3)

	other, err := s.session.GetGames(1)
	s.Nil(err)
	s.Len(other, 3)
}

func (s *dbTestingSuite) TestGetGameById() {
	game, err := s.session.GetGameById(1)
	s.Nil(err)
	s.Len(game, 3)
}

func (s *dbTestingSuite) TestUpdateGame() {
	game := models.Game{
		SeasonID:   1,
		HomeTeamID: 1,
		AwayTeamID: 1,
	}

	g, err := s.session.SaveGame(game)
	s.Nil(err)

	g.HomeTeamID = 2
	g.AwayTeamID = 2

	other, err := s.session.UpdateGame(1, *g)
	s.Nil(err)
	s.Equal(uint(1), other.ID)
	s.Equal(uint(1), other.SeasonID)
	s.Equal(uint(2), other.HomeTeamID)
	s.Equal(uint(2), other.AwayTeamID)
}

func (s *dbTestingSuite) TestUpdateGames() {
	games := []models.Game{
		{
			SeasonID:   1,
			HomeTeamID: 1,
			AwayTeamID: 1,
		},
		{
			SeasonID:   1,
			HomeTeamID: 2,
			AwayTeamID: 2,
		},
		{
			SeasonID:   1,
			HomeTeamID: 3,
			AwayTeamID: 3,
		},
	}

	g, err := s.session.SaveGames(games)
	s.Nil(err)

	for i := range g {
		g[i].HomeTeamID = 4
		g[i].AwayTeamID = 4
	}

	other, err := s.session.UpdateGames(g)
	s.Nil(err)
	s.Len(other, 3)
	for i := range other {
		s.Equal(uint(4), other[i].HomeTeamID)
		s.Equal(uint(4), other[i].AwayTeamID)
	}
}

func (s *dbTestingSuite) TestDeleteGame() {
	game := models.Game{
		SeasonID:   1,
		HomeTeamID: 1,
		AwayTeamID: 1,
	}

	g, err := s.session.SaveGame(game)
	s.Nil(err)
	s.NotNil(g)

	err = s.session.DeleteGame(1)
	s.Nil(err)

	other, err := s.session.GetGame(1)
	s.NotNil(err)
	s.Nil(other)
}

func (s *dbTestingSuite) TestDeleteGames() {
	games := []models.Game{
		{
			SeasonID:   1,
			HomeTeamID: 1,
			AwayTeamID: 1,
		},
		{
			SeasonID:   1,
			HomeTeamID: 2,
			AwayTeamID: 2,
		},
		{
			SeasonID:   1,
			HomeTeamID: 3,
			AwayTeamID: 3,
		},
	}

	g, err := s.session.SaveGames(games)
	s.Nil(err)
	s.Len(g, 3)

	err = s.session.DeleteGames(1)
	s.Nil(err)

	other, err := s.session.GetGames(1)
	s.Nil(err)
	s.Len(other, 0)
}
