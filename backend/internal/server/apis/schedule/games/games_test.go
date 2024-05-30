package games

import (
	"github.com/gofiber/fiber/v2"
	"testing"
)

func TestGenerate(t *testing.T) {
	app := fiber.New()
	app.Post("/schedule/games", handleGenerate)

	t.Run("Test handleGenerate", func(t *testing.T) {
		// TODO implement
	})

}
