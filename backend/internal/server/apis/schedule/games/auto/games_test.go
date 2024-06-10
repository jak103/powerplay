package auto

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/models"
	"github.com/stretchr/testify/assert"
	"mime/multipart"
	"net/http/httptest"
	"testing"
)

func TestGenerate(t *testing.T) {
	t.Run("Test handleOptimizeGames", func(t *testing.T) {
		// Mock Fiber app
		app := fiber.New()
		app.Put("/schedule/auto/optimize", handleOptimizeGames)

		// Make a request to the endpoint
		request := httptest.NewRequest("PUT", "/schedule/auto/optimize", nil)
		resp, err := app.Test(request)

		// Check the response
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("Test handleCreateGames", func(t *testing.T) {
		// Mock Fiber app
		app := fiber.New()
		app.Post("/schedule/auto/games", handleCreateGames)

		// Make a request to the endpoint
		request := httptest.NewRequest("POST", "/schedule/auto/games", nil)
		resp, err := app.Test(request)

		// Check the response
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("Test readBody", func(t *testing.T) {
		// Mock Fiber context
		c := &fiber.Ctx{
			Body: []byte(`{"season_id": 123, "algorithm": "round_robin", "number_of_games_per_team": 5}`),
		}

		// Call readBody function
		body, err := readBody(c)

		// Check the result
		assert.NoError(t, err)
		assert.NotNil(t, body)
		assert.Equal(t, uint(123), body.seasonID)
		assert.Equal(t, "round_robin", body.algorithm)
		assert.Equal(t, 5, body.numberOfGamesPerTeam)
		assert.Empty(t, body.iceTimes)
	})

	t.Run("Test getIceTimes", func(t *testing.T) {
		// Mock multipart file header
		file := &multipart.FileHeader{
			Filename: "test.csv",
		}

		// Call getIceTimes function
		iceTimes, err := getIceTimes(*file)

		// Check the result
		assert.Error(t, err)
		assert.Nil(t, iceTimes)
	})

	t.Run("Test assignLockerRooms", func(t *testing.T) {
		// Mock games
		games := []models.Game{
			{},
			{},
			{},
			{},
			{},
		}

		// Call assignLockerRooms function
		assignLockerRooms(games)

		// Check the result
		for _, game := range games {
			assert.NotEmpty(t, game.HomeTeamLockerRoom)
			assert.NotEmpty(t, game.AwayTeamLockerRoom)
		}
	})
}
