package stats

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/formatters"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/penalties", auth.Public, getPenaltiesHandler)
	apis.RegisterHandler(fiber.MethodGet, "/penalties/:id<int>", auth.Public, getPenaltyByIdHandler)
	apis.RegisterHandler(fiber.MethodGet, "/games/:gameID<int>/penalties", auth.Public, getPenaltiesByGameIdHandler)
	apis.RegisterHandler(fiber.MethodGet, "/teams/:teamID<int>/penalties", auth.Public, getPenaltiesByTeamIdHandler)
	apis.RegisterHandler(fiber.MethodGet, "/users/:playerID<int>/penalties", auth.Public, getPenaltiesByPlayerIdHandler)
	apis.RegisterHandler(fiber.MethodPost, "/penalties", auth.Public, postPenaltyHandler)
	apis.RegisterHandler(fiber.MethodPut, "/penalties/:id<int>", auth.Public, putPenaltyHandler)
	apis.RegisterHandler(fiber.MethodDelete, "/penalties/:id<int>", auth.Public, deletePenaltyHandler)
}

func getPenaltiesHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)
	db := db.GetSession(c)
	penalties, err := db.GetPenalties()
	if err != nil {
		log.WithErr(err).Error("Failed to get all penalties from the database")
		return responder.InternalServerError(c)
	}

	return responder.OkWithData(c, penalties)
}

func getPenaltyByIdHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	log := locals.Logger(c)
	db := db.GetSession(c)
	penalty, err := db.GetPenaltyByID(id)
	if err != nil {
		log.WithErr(err).Error("Failed to get the penalty from the database with ID %s", id)
		return responder.InternalServerError(c)
	}

	return responder.OkWithData(c, penalty)
}

func getPenaltiesByGameIdHandler(c *fiber.Ctx) error {
	gameID := c.Params("gameID")
	log := locals.Logger(c)
	db := db.GetSession(c)
	penalties, err := db.GetPenaltiesByGameID(gameID)
	if err != nil {
		log.WithErr(err).Error("Failed to get all penalties from the database for Game ID %s", gameID)
		return responder.InternalServerError(c)
	}

	return responder.OkWithData(c, penalties)
}

func getPenaltiesByTeamIdHandler(c *fiber.Ctx) error {
	teamID := c.Params("teamID")
	log := locals.Logger(c)
	db := db.GetSession(c)
	penalties, err := db.GetPenaltiesByTeamID(teamID)
	if err != nil {
		log.WithErr(err).Error("Failed to get all penalties from the database for Team ID %s", teamID)
		return responder.InternalServerError(c)
	}

	return responder.OkWithData(c, penalties)
}

func getPenaltiesByPlayerIdHandler(c *fiber.Ctx) error {
	playerID := c.Params("playerID")
	log := locals.Logger(c)
	db := db.GetSession(c)
	penalties, err := db.GetPenaltiesByPlayerID(playerID)
	if err != nil {
		log.WithErr(err).Error("Failed to get all penalties from the database for Player ID %s", playerID)
		return responder.InternalServerError(c)
	}

	return responder.OkWithData(c, penalties)
}

func postPenaltyHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)

	// Parse penalty
	penaltyRequest := &models.Penalty{}
	err := c.BodyParser(penaltyRequest)
	if err != nil {
		log.WithErr(err).Error("Failed to parse penalty request payload")
		return responder.BadRequest(c, "Failed to parse penalty request payload")
	}

	// Validate request
	validate := validator.New()
	err = validate.Struct(penaltyRequest)
	if err != nil {
		return responder.BadRequest(c, "Failed to validate request")
	}

	// Create penalty
	db := db.GetSession(c)
	penalty, err := db.CreatePenalty(penaltyRequest)
	if err != nil {
		log.WithErr(err).Error("Failed to save penalty request")
		return responder.InternalServerError(c)
	}

	return responder.OkWithData(c, penalty)
}

func putPenaltyHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	log := locals.Logger(c)

	// Parse penalty
	penaltyRequest := &models.Penalty{}
	err := c.BodyParser(penaltyRequest)
	if err != nil {
		log.WithErr(err).Error("Failed to parse penalty request payload")
		return responder.BadRequest(c, "Failed to parse penalty request payload")
	}
	if formatters.UintToString(penaltyRequest.ID) != id {
		log.WithErr(err).Error("Penalty ID in URL does not match ID in payload")
		return responder.BadRequest(c, "Penalty ID in URL does not match ID in payload")
	}

	// Validate request
	validate := validator.New()
	err = validate.Struct(penaltyRequest)
	if err != nil {
		return responder.BadRequest(c, "Failed to validate request")
	}

	// Update penalty
	db := db.GetSession(c)
	penalty, err := db.UpdatePenalty(penaltyRequest)
	if err != nil {
		log.WithErr(err).Error("Failed to update penalty with ID %s", id)
		return responder.InternalServerError(c)
	}

	return responder.OkWithData(c, penalty)
}

func deletePenaltyHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	log := locals.Logger(c)
	db := db.GetSession(c)

	// Get penalty
	penalty, err := db.GetPenaltyByID(id)
	if penalty == nil {
		log.WithErr(err).Error("Delete penalty failed because there is no penalty with ID %s", id)
		return responder.BadRequest(c, "Delete penalty failed because there is no penalty with ID %s", id)
	}

	// Delete penalty
	// Yes this could just be done first, but checking that the penalty exists allows for a more specific error message
	err = db.DeletePenalty(penalty)
	if err != nil {
		log.WithErr(err).Error("Failed to delete penalty request")
		return responder.InternalServerError(c)
	}

	return responder.Ok(c)
}
