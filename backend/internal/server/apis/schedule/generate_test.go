package schedule

import (
	"bytes"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis/schedule/pkg/models"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {

	t.Run("Test HandleGenerate", func(t *testing.T) {
		app := fiber.New()
		body := `{"seasonFileName":"test", "numberOfGamesPerTeam": 10}`

		req := httptest.NewRequest("POST", "/generate", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("Test ReadBody", func(t *testing.T) {
		c := &fiber.Ctx{}
		body := `{"seasonFileName":"test", "numberOfGamesPerTeam": 10}`
		c.Request().SetBody([]byte(body))

		seasonFileName, numberOfGamesPerTeam, err := ReadBody(c)
		assert.Nil(t, err)
		assert.Equal(t, "season.json", seasonFileName)
		assert.Equal(t, 10, numberOfGamesPerTeam)
	})

	t.Run("Test OptimizeSchedule", func(t *testing.T) {
		games := []models.Game{
			{Start: time.Now()},
			{Start: time.Now().Add(1 * time.Hour)},
		}

		OptimizeSchedule(games)

		assert.NotEmpty(t, games)
	})

	t.Run("Test GenerateGames", func(t *testing.T) {
		leagues := []models.League{
			{Name: "League1", Teams: []models.Team{{Id: "1", Name: "Team1"}, {Id: "2", Name: "Team2"}}},
		}

		season := GenerateGames(leagues, 2)

		assert.NotEmpty(t, season.LeagueRounds)
		assert.Equal(t, 2, len(season.LeagueRounds["League1"][0].Games))
	})

	t.Run("Test AssignTimes", func(t *testing.T) {
		times := []string{"1/2/23 20:00", "1/3/23 21:00"}
		season := models.Season{
			LeagueRounds: map[string][]models.Round{
				"League1": {
					{Games: []models.Game{{}, {}}},
				},
			},
		}

		games := AssignTimes(times, season)

		assert.Equal(t, 2, len(games))
		assert.Equal(t, "20:00", games[0].StartTime)
	})

	t.Run("Test GetBalanceCount", func(t *testing.T) {
		teamStats := map[string]models.TeamStats{
			"Team1": {Balanced: true},
			"Team2": {Balanced: false},
		}

		count := GetBalanceCount(&teamStats)

		assert.Equal(t, 1, count)
	})

	t.Run("Test RotateTeams", func(t *testing.T) {
		league := models.League{
			Teams: []models.Team{{Id: "1", Name: "Team1"}, {Id: "2", Name: "Team2"}, {Id: "3", Name: "Team3"}},
		}

		RotateTeams(&league)

		assert.Equal(t, "Team1", league.Teams[0].Name)
		assert.Equal(t, "Team3", league.Teams[1].Name)
		assert.Equal(t, "Team2", league.Teams[2].Name)
	})

	t.Run("Test NewGame", func(t *testing.T) {
		game := NewGame("League1", "1", "Team1", "2", "Team2")

		assert.Equal(t, "Team1", game.Team1Name)
		assert.Equal(t, "Team2", game.Team2Name)
		assert.Equal(t, "League1", game.League)
	})

	t.Run("Test NewGames", func(t *testing.T) {
		season := models.Season{
			LeagueRounds: map[string][]models.Round{
				"League1": {
					{Games: []models.Game{{Team1Id: "1", Team2Id: "2"}, {Team1Id: "-1", Team2Id: "3"}}},
				},
			},
		}

		games := NewGames(&season)

		assert.Equal(t, 1, len(games)) // Only one game should be added (the game without a bye)
		assert.Equal(t, "1", games[0].Team1Id)
	})

	t.Run("Test IsEarlyGame", func(t *testing.T) {
		assert.True(t, IsEarlyGame(20, 0))
		assert.True(t, IsEarlyGame(21, 15))
		assert.False(t, IsEarlyGame(21, 30))
		assert.False(t, IsEarlyGame(22, 0))
	})
}
