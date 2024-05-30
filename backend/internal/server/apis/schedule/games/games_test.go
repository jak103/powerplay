package games

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestGenerate(t *testing.T) {
	app := fiber.New()
	app.Post("/schedule/games", handleGenerate)

	t.Run("Test handleGenerate", func(t *testing.T) {
		body := `{"seasonName":"test", "numberOfGamesPerTeam": 10}`

		req := httptest.NewRequest("POST", "/schedule/games", bytes.NewBufferString(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := app.Test(req)

		assert.Nil(t, err)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

}
