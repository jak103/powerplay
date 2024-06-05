package manual

import (
	"github.com/gofiber/fiber/v2"
	"testing"
)

func TestGenerate(t *testing.T) {
	app := fiber.New()
	app.Post("/schedule/games", handleCreateGames)

	t.Run("Test handleCreateGames", func(t *testing.T) {
		// TODO implement
	})
}
