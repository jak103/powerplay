package components

import (
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/stretchr/testify/assert"
)

func TestPostSeasonBadRequest(t *testing.T) {
	app := apis.CreateTestApp()

	req, _ := http.NewRequest(fiber.MethodPost, "/api/v1/seasons", nil)

	res, err := app.Test(req)

	assert.Nil(t, err)

	assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
}

func TestPostGoal(t *testing.T) {
	database.newSession = func(c *fiber.Ctx) session {
		return &sessionMock{
			CreateSeasonFunc: func(season *models.Season) (*models.Season, error) {
				return &models.Season{
						Name: "Season 1",
					},
					nil
			},
		}
	}

	bodyReader := strings.NewReader(`{"name": "Season 1","start": "2023-05-01T00:00:00Z","end": "2023-09-01T00:00:00Z"}`)

	app := apis.CreateTestApp()

	req, err := http.NewRequest(fiber.MethodPost, "/api/v1/seasons", bodyReader)
	assert.Nil(t, err)
	req.Header.Add("Content-Type", "application/json")

	res, err := app.Test(req)

	assert.Nil(t, err)

	assert.Equal(t, fiber.StatusOK, res.StatusCode)
}

func TestGetGoal(t *testing.T) {
	database.newSession = func(c *fiber.Ctx) session {
		return &sessionMock{
			GetSeasonsFunc: func() ([]models.Season, error) {
				return []models.Season{{Name: "Season1"}, {Name: "Season2"}}, nil

			},
		}
	}

	app := apis.CreateTestApp()

	req, err := http.NewRequest(fiber.MethodGet, "/api/v1/seasons", nil)
	assert.Nil(t, err)
	req.Header.Add("Content-Type", "application/json")

	res, err := app.Test(req)

	assert.Nil(t, err)

	assert.Equal(t, fiber.StatusOK, res.StatusCode)
}
