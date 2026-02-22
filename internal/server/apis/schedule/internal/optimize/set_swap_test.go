package optimize

import (
	"testing"
	"time"

	"github.com/jak103/powerplay/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestOptimizeSchedule(t *testing.T) {
	// create a fake schedule where team 0 has a bunch of early games and team 1 has bunch of late games
	var games = []models.Game{
		models.Game{
			HomeTeam: models.Team{
				Name: "team0",
				League: models.League{
					Name: "league0",
				},
			},
			AwayTeam: models.Team{
				Name: "team3",
				League: models.League{
					Name: "league0",
				},
			},
			Start: time.Date(2024, 6, 5, 20, 0, 0, 0, time.UTC),
		},
		models.Game{
			HomeTeam: models.Team{
				Name: "team0",
				League: models.League{
					Name: "league0",
				},
			},
			AwayTeam: models.Team{
				Name: "team4",
				League: models.League{
					Name: "league0",
				},
			},
			Start: time.Date(2024, 6, 10, 20, 0, 0, 0, time.UTC),
		},
		models.Game{
			HomeTeam: models.Team{
				Name: "team0",
				League: models.League{
					Name: "league0",
				},
			},
			AwayTeam: models.Team{
				Name: "team5",
				League: models.League{
					Name: "league0",
				},
			},
			Start: time.Date(2024, 6, 15, 20, 0, 0, 0, time.UTC),
		},
		models.Game{
			HomeTeam: models.Team{
				Name: "team1",
				League: models.League{
					Name: "league0",
				},
			},
			AwayTeam: models.Team{
				Name: "team3",
				League: models.League{
					Name: "league0",
				},
			},
			Start: time.Date(2024, 6, 5, 22, 0, 0, 0, time.UTC),
		},
		models.Game{
			HomeTeam: models.Team{
				Name: "team1",
				League: models.League{
					Name: "league0",
				},
			},
			AwayTeam: models.Team{
				Name: "team4",
				League: models.League{
					Name: "league0",
				},
			},
			Start: time.Date(2024, 6, 10, 22, 0, 0, 0, time.UTC),
		},
	}

	SetOptimizeSchedule(games)

	assert.Equal(t, "team1", games[0].HomeTeam.Name)
}
