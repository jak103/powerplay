package components

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/teams", auth.Public, getTeams)
	apis.RegisterHandler(fiber.MethodGet, "/teams/:teamId<int>", auth.Public, getTeam)
	apis.RegisterHandler(fiber.MethodPut, "/teams/:teamId<int>", auth.Public, updateTeam)
	apis.RegisterHandler(fiber.MethodPost, "/teams", auth.Public, createTeam)
}

// Handler to get all Team
func getTeams(c *fiber.Ctx) error {
	log := locals.Logger(c)
	db := db.GetSession(c)
	teams, err := db.GetTeams()
	if err != nil {
		log.WithErr(err).Alert("Failed to get all Teams from the database")
		return responder.InternalServerError(c)
	}
	return responder.OkWithData(c, teams)
}

// Handler to get team details by ID
func getTeam(c *fiber.Ctx) error {
	teamID := c.Params("teamId")
	log := locals.Logger(c)
	db := db.GetSession(c)
	team, err := db.GetTeamByID(teamID)
	if err != nil {
		log.WithErr(err).Alert("Failed to get the Team from the database")
		return responder.InternalServerError(c)
	}
	return responder.OkWithData(c, team)
}

// Handler to update team details by ID
func updateTeam(c *fiber.Ctx) error {
	teamID := c.Params("teamId")
	log := locals.Logger(c)
	db := db.GetSession(c)
	var team models.Team
	if err := c.BodyParser(&team); err != nil {
		log.WithErr(err).Alert("Failed to parse team data")
		return responder.BadRequest(c, "Failed to parse team data")
	}
	if err := db.UpdateTeam(teamID, team); err != nil {
		log.WithErr(err).Alert("Failed to update Team in the database")
		return responder.InternalServerError(c)
	}
	return responder.Ok(c)
}

// Handler to create a new team
func createTeam(c *fiber.Ctx) error {
	db := db.GetSession(c)
	log := locals.Logger(c)
	var team models.Team
	if err := c.BodyParser(&team); err != nil {
		log.WithErr(err).Alert("Failed to parse team data")
		return responder.BadRequest(c, "Failed to parse team data")
	}

	newTeam, err := db.CreateTeam(&team)
	if err != nil {
		log.WithErr(err).Alert("Failed to create Team in the database")
		return responder.InternalServerError(c)
	}
	return responder.OkWithData(c, newTeam) // TODO: 201 Created and return the team ID
}
