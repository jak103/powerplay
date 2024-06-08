package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/games", auth.Public, getGames)
	apis.RegisterHandler(fiber.MethodGet, "/games/:gameId", auth.Public, getGame)
	apis.RegisterHandler(fiber.MethodPut, "/games/:gameId", auth.Public, updateGame)
	apis.RegisterHandler(fiber.MethodPost, "/games", auth.Public, createGame)
}

// Handler to get all games
func getGames(c *fiber.Ctx) error {
	db := db.GetSession(c)
	games, err := db.GetGames()
	if err != nil {
		log.WithErr(err).Alert("Failed to get all Games from the database")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return responder.OkWithData(c, games)
}

// Handler to get game details by ID
func getGame(c *fiber.Ctx) error {
	gameID := c.Params("gameId")
	db := db.GetSession(c)
	game, err := db.GetGameByID(gameID)
	if err != nil {
		log.WithErr(err).Alert("Failed to get Game from the database")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return responder.OkWithData(c, game)
}

// Handler to update game details by ID
func updateGame(c *fiber.Ctx) error {
	gameID := c.Params("gameId")
	db := db.GetSession(c)
	var game models.Game
	if err := c.BodyParser(&game); err != nil {
		log.WithErr(err).Alert("Failed to parse game data")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := db.UpdateGame(gameID, game); err != nil {
		log.WithErr(err).Alert("Failed to update Game in the database")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusOK)
}

// Handler to create a new game
func createGame(c *fiber.Ctx) error {
	db := db.GetSession(c)
	var game models.Game
	if err := c.BodyParser(&game); err != nil {
		log.WithErr(err).Alert("Failed to parse game data")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	if err := db.CreateGame(&game); err != nil {
		log.WithErr(err).Alert("Failed to create Game in the database")
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusCreated)
}
