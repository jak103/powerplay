package users

import (
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/utils/unittesting"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	newSession = func(_ *fiber.Ctx) session {
		return &sessionMock{
			CreateUserFunc: func(email, passwordHash, salt string) error {
				return nil
			},
			GetUserByEmailFunc: func(email string) (*models.User, error) {
				return nil, nil
			},
		}
	}

	ctx := unittesting.FiberCtx()
	ctx.Request().Header.Set("Content-Type", "application/json")
	ctx.Request().SetBody([]byte(`{ "email": "awesome_skater@email.com", "password": "midnight9"}`))
	err := createUserAccount(ctx)
	assert.NoError(t, err)

	assert.Equal(t, fiber.StatusOK, ctx.Response().StatusCode())
}
