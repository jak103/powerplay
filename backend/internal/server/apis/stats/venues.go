package stats

import (
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/utils/responder"
	"github.com/jak103/powerplay/internal/server/services/auth"
)

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/venues", auth.Public, postVenueHandler)
	apis.RegisterHandler(fiber.MethodGet, "/venues", auth.Public, getVenuesHandler)
}

func postVenueHandler (c *fiber.Ctx) error {
	log := locals.Logger(c)
	log.Debug("body: %q", c.Request().Body())
	venueRequest := &models.Venue{}

	err := c.BodyParser(venueRequest)
	if err != nil {
		log.WithErr(err).Error("Failed to parse venue request payload")
		return responder.BadRequest(c, "Failed to parse venue request payload")
	}
	db := db.GetSession(c)
	record, err := db.SaveVenue(venueRequest)
	
	if err != nil {
		log.WithErr(err).Error("Failed to save venue request")
		return responder.InternalServerError(c)
	}

	if record == nil {
		return responder.BadRequest(c, "Could not post venue into database")
	}
	return responder.Ok(c)
}

// todo finish getVenuesHandler
func getVenuesHandler (c *fiber.Ctx) error {
	log := locals.Logger(c)
	db := db.GetSession(c)
	venues, err := db.GetVenues()
	
	if err != nil {
		log.WithErr(err).Error("Failed to get venues from the database ")
	}

	return responder.OkWithData(c, venues)
}