package team

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/log"
)

func init() {

	apis.RegisterHandler(fiber.MethodGet, "/Team/:teamId", auth.Public, getTeam)
	apis.RegisterHandler(fiber.MethodPut, "/Team/:teamId", auth.Public, updateTeam)
	apis.RegisterHandler(fiber.MethodPost, "/Team", auth.Public, createTeam)
}

// Handler to get all Team
func getTeams(c *fiber.Ctx) error {
	db := db.GetSession(c)
	Team, err := db.GetTeam()
	if err != nil {
		log.WithErr(err).Alert("Failed to get all Team from the database")
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(Team)
}

// Handler to get team details by ID
func getTeam(c *fiber.Ctx) error {
	teamID := c.Params("teamId")
	db := db.GetSession(c)
	team, err := db.GetTeamByID(teamID)
	if err != nil {
		log.WithErr(err).Alert("Failed to get Team from the database")
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(team)
}

// Handler to update team details by ID
func updateTeam(c *fiber.Ctx) error {
	teamID := c.Params("teamId")
	db := db.GetSession(c)
	var team models.Team
	if err := c.BodyParser(&team); err != nil {
		log.WithErr(err).Alert("Failed to parse team data")
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := db.UpdateTeam(teamID, team); err != nil {
		log.WithErr(err).Alert("Failed to update Team in the database")
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}

// Handler to create a new team
func createTeam(c *fiber.Ctx) error {
	db := db.GetSession(c)
	var team models.Team
	if err := c.BodyParser(&team); err != nil {
		log.WithErr(err).Alert("Failed to parse team data")
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if err := db.CreateTeam(&team); err != nil {
		log.WithErr(err).Alert("Failed to create Team in the database")
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusCreated)
}
