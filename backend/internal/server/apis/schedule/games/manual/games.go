package manual

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
	apis.RegisterHandler(fiber.MethodPost, "/schedule/manual/game", auth.Authenticated, handleCreateGame)
	apis.RegisterHandler(fiber.MethodPost, "/schedule/manual/games", auth.Authenticated, handleCreateGames)
	apis.RegisterHandler(fiber.MethodGet, "/schedule/manual/games/:id", auth.Authenticated, handleGetGame)
	apis.RegisterHandler(fiber.MethodGet, "/schedule/manual/games", auth.Authenticated, handleGetGames)
	apis.RegisterHandler(fiber.MethodPut, "/schedule/manual/games/:id", auth.Authenticated, handleUpdateGame)
	apis.RegisterHandler(fiber.MethodPut, "/schedule/manual/games", auth.Authenticated, handleUpdateGames)
	apis.RegisterHandler(fiber.MethodDelete, "/schedule/manual/games/:id", auth.Authenticated, handleDeleteGame)
	apis.RegisterHandler(fiber.MethodDelete, "/schedule/manual/games", auth.Authenticated, handleDeleteGames)
}

// basic CRUD operations
// - the caller is allowed to fine tune the schedule by updating the games

func handleCreateGame(c *fiber.Ctx) error {
	type Dto struct {
		Game models.Game `json:"game"`
	}
	var dto Dto
	err := c.BodyParser(&dto)
	if err != nil {
		log.Error("Failed to parse request body")
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	game := dto.Game
	session := db.GetSession(c)
	_, err = session.SaveGame(game)
	if err != nil {
		log.Error("Failed to save game to the database")
		return responder.InternalServerError(c, err)
	}
	return responder.Ok(c, game)
}

func handleCreateGames(c *fiber.Ctx) error {
	type Dto struct {
		Games []models.Game `json:"games"`
	}
	var dto Dto
	err := c.BodyParser(&dto)
	if err != nil {
		log.Error("Failed to parse request body")
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	games := dto.Games
	session := db.GetSession(c)
	_, err = session.SaveGames(games)
	if err != nil {
		log.Error("Failed to save games to the database")
		return responder.InternalServerError(c, err)
	}
	return responder.Ok(c, games)
}

func handleGetGame(c *fiber.Ctx) error {
	type Dto struct {
		ID uint `json:"id"`
	}
	var dto Dto
	err := c.BodyParser(&dto)
	if err != nil {
		log.Error("Failed to parse request body")
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	id := dto.ID
	session := db.GetSession(c)
	game, err := session.GetGame(id)
	if err != nil {
		log.Error("Failed to get game from the database")
		return responder.InternalServerError(c, err)
	}
	return responder.OkWithData(c, game)
}

func handleGetGames(c *fiber.Ctx) error {
	type Dto struct {
		SeasonID uint `json:"season_id"`
	}
	var dto Dto
	err := c.BodyParser(&dto)
	if err != nil {
		log.Error("Failed to parse request body")
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	seasonID := dto.SeasonID
	session := db.GetSession(c)
	games, err := session.GetGames(seasonID)
	if err != nil {
		log.Error("Failed to get games from the database")
		return responder.InternalServerError(c, err)
	}
	return responder.OkWithData(c, games)
}

func handleUpdateGame(c *fiber.Ctx) error {
	type Dto struct {
		Game models.Game `json:"game"`
	}
	var dto Dto
	err := c.BodyParser(&dto)
	if err != nil {
		log.Error("Failed to parse request body")
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	game := dto.Game
	session := db.GetSession(c)
	_, err = session.UpdateGame(game)
	if err != nil {
		log.Error("Failed to update game in the database")
		return responder.InternalServerError(c, err)
	}
	return responder.Ok(c, game)
}

func handleUpdateGames(c *fiber.Ctx) error {
	type Dto struct {
		Games []models.Game `json:"games"`
	}
	var dto Dto
	err := c.BodyParser(&dto)
	if err != nil {
		log.Error("Failed to parse request body")
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	games := dto.Games
	session := db.GetSession(c)
	_, err = session.UpdateGames(games)
	if err != nil {
		log.Error("Failed to update games in the database")
		return responder.InternalServerError(c, err)
	}
	return responder.Ok(c, games)
}

func handleDeleteGame(c *fiber.Ctx) error {
	type Dto struct {
		ID uint `json:"id"`
	}
	var dto Dto
	err := c.BodyParser(&dto)
	if err != nil {
		log.Error("Failed to parse request body")
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	id := dto.ID
	session := db.GetSession(c)
	err = session.DeleteGame(id)
	if err != nil {
		log.Error("Failed to delete game from the database")
		return responder.InternalServerError(c, err)
	}
	return responder.Ok(c)
}

func handleDeleteGames(c *fiber.Ctx) error {
	type Dto struct {
		SeasonID uint `json:"season_id"`
	}
	var dto Dto
	err := c.BodyParser(&dto)
	if err != nil {
		log.Error("Failed to parse request body")
		return responder.BadRequest(c, fiber.StatusBadRequest, err.Error())
	}
	seasonID := dto.SeasonID
	session := db.GetSession(c)
	err = session.DeleteGames(seasonID)
	if err != nil {
		log.Error("Failed to delete games from the database")
		return responder.InternalServerError(c, err)
	}
	return responder.Ok(c)
}
