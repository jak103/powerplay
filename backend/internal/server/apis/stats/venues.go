package stats

import (
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
)

func init() {
	apis.RegisterHandler(fiber.MethodPost, "/venues", auth.Public, postVenueHandler)
}

func postVenueHandler (c *fiber.Ctx) error {
	log := locals.Logger(c)
	log.Debug("body: %q", c.Request().Body())
	venueRequest := &models.Venue{}

	err := c.BodyParser(venueRequest)
	if err != nil {
		log.WithErr(err).Alert("Failed to parse venue request payload")
		return responder.BadRequest(c, "Failed to parse venue request payload")
	}
	db := db.GetSession(c)
	err = db.CreateVenue(venueRequest)
	
	if err != nil {
		log.WithErr(err).Alert("Failed to save venue request")
		return responder.InternalServerError(c)
	}
	return responder.Ok(c)
}

// todo finish getVenuesHandler
func getVenuesHandler (c *fiber.Ctx) error P {
	log := locals.Logger(c)
	db := db.GetSession(c)
	venues, err := db.GetVenues()
	
	if err != nil {
		log.WithErr(err).Alert("Failed to get venues from the database ")
	}

	return responder.OkWithData(c, venues)
}