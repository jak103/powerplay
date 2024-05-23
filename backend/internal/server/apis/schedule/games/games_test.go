package games

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/models"
	pkgModels "github.com/jak103/powerplay/internal/server/apis/schedule/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGenerate(t *testing.T) {
	app := fiber.New()
	app.Post("/schedule/games", handleGames)

	t.Run("Test isEarlyGame", func(t *testing.T) {
		assert.True(t, isEarlyGame(20, 0))
		assert.True(t, isEarlyGame(21, 15))
		assert.False(t, isEarlyGame(21, 30))
		assert.False(t, isEarlyGame(22, 0))
	})

	t.Run("Test rotateTeams", func(t *testing.T) {
		league := models.League{
			Teams: []pkgModels.Team{{Id: "1", Name: "Team1"}, {Id: "2", Name: "Team2"}, {Id: "3", Name: "Team3"}},
		}

		rotateTeams(&league)

		assert.Equal(t, "Team1", league.Teams[0].Name)
		assert.Equal(t, "Team3", league.Teams[1].Name)
		assert.Equal(t, "Team2", league.Teams[2].Name)
	})

	t.Run("Test newGame", func(t *testing.T) {
		game := newGame("League1", "1", "Team1", "2", "Team2")

		assert.Equal(t, "Team1", game.Team1Name)
		assert.Equal(t, "Team2", game.Team2Name)
		assert.Equal(t, "League1", game.League)
	})

	t.Run("Test newGames", func(t *testing.T) {
		season := pkgModels.Season{
			LeagueRounds: map[string][]pkgModels.Round{
				"A": {
					{Games: []pkgModels.Game{{Team1Id: "1", Team2Id: "2"}, {Team1Id: "-1", Team2Id: "3"}}},
				},
			},
		}

		games, err := newGames(&season, 10)
		if err != nil {
			t.Errorf("newGames() error = %v", err)
			return
		}

		assert.Equal(t, 1, len(games)) // Only one games should be added (the games without a bye)
		assert.Equal(t, "1", games[0].Team1Id)
	})

	t.Run("Test getBalanceCount", func(t *testing.T) {
		teamStats := map[string]pkgModels.TeamStats{
			"Team1": {Balanced: true},
			"Team2": {Balanced: false},
		}

		count := getBalanceCount(&teamStats)

		assert.Equal(t, 1, count)
	})

	t.Run("Test assignTimes", func(t *testing.T) {
		times := []string{"1/2/23 20:00", "1/3/23 21:00"}

		season := pkgModels.Season{
			LeagueRounds: map[string][]pkgModels.Round{
				"A": {
					{Games: []pkgModels.Game{{Team1Id: "1", Team2Id: "2"}, {Team1Id: "3", Team2Id: "4"}}},
				},
			},
		}

		games, err := assignTimes(times, season, 2)
		if err != nil {
			t.Errorf("assignTimes() error = %v", err)
			return
		}

		assert.Equal(t, 2, len(games))
		assert.Equal(t, "20:00", games[0].StartTime)
		assert.Equal(t, "01/02/2023", games[0].StartDate) // Added assertion for StartDate
		assert.Equal(t, "21:00", games[1].StartTime)
		assert.Equal(t, "01/03/2023", games[1].StartDate) // Added assertion for StartDate
	})

	t.Run("Test optimizeSchedule", func(t *testing.T) {
		games := []pkgModels.Game{
			{Start: time.Now()},
			{Start: time.Now().Add(1 * time.Hour)},
		}

		optimizeSchedule(games)

		assert.NotEmpty(t, games)
	})

	t.Run("Test generateGames", func(t *testing.T) {
		leagues := []models.League{
			{Name: "League1", Teams: []pkgModels.Team{{Id: "1", Name: "Team1"}, {Id: "2", Name: "Team2"}}},
		}

		season, err := generateGames(leagues, 2)
		if err != nil {
			t.Errorf("generateGames() error = %v", err)
			return
		}

		assert.NotEmpty(t, season.LeagueRounds)
		assert.Equal(t, 1, len(season.LeagueRounds["League1"][0].Games))
	})

	t.Run("Test readBody", func(t *testing.T) {
		app := fiber.New()

		c := app.AcquireCtx(&fasthttp.RequestCtx{})
		defer app.ReleaseCtx(c)
		body := `{"seasonName":"test", "numberOfGamesPerTeam": 10}`
		c.Request().SetBody([]byte(body))

		seasonFileName, numberOfGamesPerTeam, err := readBody(c)
		assert.Nil(t, err)
		assert.Equal(t, "test", seasonFileName)
		assert.Equal(t, 10, numberOfGamesPerTeam)
	})

	t.Run("Test handleGames", func(t *testing.T) {
		body := `{"seasonName":"test", "numberOfGamesPerTeam": 10}`

		req := httptest.NewRequest("POST", "/schedule/games", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

}
