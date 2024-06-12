package stats

import (
	"net/http"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/stretchr/testify/assert"
)

func TestPostGoalBadRequest(t *testing.T) {
	app := apis.CreateTestApp()

	req, _ := http.NewRequest(fiber.MethodPost, "/api/v1/goals", nil)

	res, err := app.Test(req)

	assert.Nil(t, err)

	assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
}

func TestPostGoal(t *testing.T) {
	database.newSession = func(c *fiber.Ctx) session {
		return &sessionMock{
			SaveGoalFunc: func(goal *models.Goal) (*models.Goal, error) {
				return &models.Goal{
						UserId: 1,
					},
					nil
			},
		}
	}

	bodyReader := strings.NewReader(`{"user_id": 12,"game_id": 13,"team_id": 1,"duration": 27,"period": 3,"assist1_id": 57,"assist2_id": 84,"playerdifferential": -2,"ispenaltyshot": true }`)

	app := apis.CreateTestApp()

	req, _ := http.NewRequest(fiber.MethodPost, "/api/v1/goals", bodyReader)
	req.Header.Add("Content-Type", "application/json")

	res, err := app.Test(req)

	assert.Nil(t, err)

	assert.Equal(t, fiber.StatusOK, res.StatusCode)
}
