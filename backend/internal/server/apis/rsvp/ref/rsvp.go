package ref

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
	apis.RegisterHandler(fiber.MethodPut, "/rsvp/ref", auth.Authenticated, assignRefToGame)
}

func assignRefToGame(c *fiber.Ctx) error {
	type BodyDto struct {
		RefID     uint `json:"refId"`
		GameID    uint `json:"gameId"`
		IsPrimary bool `json:"isPrimary"`
	}

	var bodyDto BodyDto
	err := c.BodyParser(bodyDto)
	if err != nil {
		return responder.BadRequest(c, "Error reading body of assignRefToGame: %v", err)
	}

	refId, gameId, isPrimary, err := bodyDto.RefID, bodyDto.GameID, bodyDto.IsPrimary, nil

	// Get the ref from the db
	s := db.GetSession(nil)
	refUser, err := s.GetRefById(refId)

	if err != nil {
		return responder.BadRequest(c, "Error getting ref user from the db: %v", err)
	}

	// Get the game fro mthe db
	game, err := s.GetGame(gameId)
	if err != nil {
		return responder.BadRequest(c, "Error getting game from the db: %v", err)
	}

	// Assign the ref to the game
	if game != nil && refUser != nil {
		if isPrimary {
			game.PrimaryReferee = refUser
			game.PrimaryRefereeID = &refUser.ID
		} else {
			game.SecondaryReferee = refUser
			game.SecondaryRefereeID = &refUser.ID
		}
	}

	// Save the update
	log.Debug("Saving referee: %v sign up to the game %v\n", refId, gameId)
	gameList := make([]models.Game, 0)
	gameList = append(gameList, *game)

	returnedGames, err := s.SaveGames(gameList)
	if err != nil {
		log.Error("Error saving game update to database: %v\n", err)
		return responder.InternalServerError(c)
	}

	return responder.OkWithData(c, returnedGames)
}
