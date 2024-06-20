package stats

import (
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
	apis.RegisterHandler(fiber.MethodGet, "/penaltyTypes", auth.Public, getPenaltyTypes)
	apis.RegisterHandler(fiber.MethodPost, "/penaltyTypes", auth.Public, postPenaltyTypeHandler)
	apis.RegisterHandler(fiber.MethodPut, "/penaltyTypes/:id<int>", auth.Public, putPenaltyTypeHandler)
	apis.RegisterHandler(fiber.MethodDelete, "/penaltyTypes/:id<int>", auth.Public, deletePenaltyTypeHandler)
}

func getPenaltyTypes(c *fiber.Ctx) error {
	log := locals.Logger(c)
	db := db.GetSession(c)
	penaltyTypes, err := db.GetPenaltyTypes()
	if err != nil {
		log.WithErr(err).Error("Failed to get all penalty types from the database")
		return responder.InternalServerError(c)
	}

	return responder.OkWithData(c, penaltyTypes)
}

func postPenaltyTypeHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)

	// Parse penalty type
	penaltyType := &models.PenaltyType{}
	err := c.BodyParser(penaltyType)
	if err != nil {
		log.WithErr(err).Error("Failed to parse penalty type request payload")
		return responder.BadRequest(c, "Failed to parse penalty type request payload")
	}

	db := db.GetSession(c)
	penaltyType, err = db.CreatePenaltyType(penaltyType)
	if err != nil {
		log.WithErr(err).Error("Failed to save penalty type request")
		return responder.InternalServerError(c)
	}

	return responder.OkWithData(c, penaltyType)
}

func putPenaltyTypeHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	log := locals.Logger(c)

	// Parse penalty type
	penaltyType := &models.PenaltyType{}
	err := c.BodyParser(penaltyType)
	if err != nil {
		log.WithErr(err).Error("Failed to parse penalty type request payload")
		return responder.BadRequest(c, "Failed to parse penalty type request payload")
	}
	if formatters.UintToString(penaltyType.ID) != id {
		log.WithErr(err).Error("Penalty Type ID in URL does not match ID in payload")
		return responder.BadRequest(c, "Penalty Type ID in URL does not match ID in payload")
	}

	// Update penalty type
	db := db.GetSession(c)
	newPenaltyType, err := db.UpdatePenaltyType(penaltyType)
	if err != nil {
		log.WithErr(err).Error("Failed to update penalty type with ID %s", id)
		return responder.InternalServerError(c)
	}

	return responder.OkWithData(c, newPenaltyType)
}

func deletePenaltyTypeHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	log := locals.Logger(c)
	db := db.GetSession(c)

	// Get penalty type
	penaltyType, err := db.GetPenaltyTypeByID(id)
	if penaltyType == nil {
		log.WithErr(err).Error("Delete penalty type failed because there is no penalty type with ID %s", id)
		return responder.BadRequest(c, "Delete penalty type failed because there is no penalty type with ID %s", id)
	}

	// Delete penalty type
	// Yes this could just be done first, but checking that the penalty type exists allows for a more specific error message
	err = db.DeletePenaltyType(penaltyType)
	if err != nil {
		log.WithErr(err).Error("Failed to delete penalty type request")
		return responder.InternalServerError(c)
	}

	return responder.Ok(c)
}
