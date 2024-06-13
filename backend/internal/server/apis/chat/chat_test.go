package chat

import (
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/stretchr/testify/assert"
)

func TestChat(t *testing.T) {
	app := apis.CreateTestApp()

	req, _ := http.NewRequest(fiber.MethodGet, "/api/v1/hello", nil)

	res, err := app.Test(req)

	assert.Nil(t, err)

	assert.Equal(t, fiber.StatusOK, res.StatusCode)
}
