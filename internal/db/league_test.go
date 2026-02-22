package db

import "github.com/jak103/powerplay/internal/models"

func (s *dbTestingSuite) TestGetLeaguesBySeasonId() {
	season := models.Season{
		Name: "Season 1",
	}

	dbSeason, err := s.session.CreateSeason(&season)
	s.Nil(err)

	league := models.League{
		Name:     "League 1",
		SeasonID: dbSeason.ID,
	}

	dbLeague, err := s.session.CreateLeague(&league)
	s.Nil(err)
	s.NotNil(dbLeague)

	team1 := models.Team{
		Name:     "Team 1",
		LeagueID: dbLeague.ID,
	}

	team2 := models.Team{
		Name:     "Team2",
		LeagueID: dbLeague.ID,
	}

	_, err = s.session.CreateTeam(&team1)
	s.Nil(err)

	_, err = s.session.CreateTeam(&team2)
	s.Nil(err)

	leagues, err := s.session.GetLeaguesBySeason(int(dbSeason.ID))
	s.Nil(err)
	s.Equal(1, len(leagues))
	s.Equal(2, len(leagues[0].Teams))
	s.Equal("Team 1", leagues[0].Teams[0].Name)
}
