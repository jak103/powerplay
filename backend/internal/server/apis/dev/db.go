package admin

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/log"
	"github.com/jak103/powerplay/internal/utils/responder"
)

func init() {
	apis.RegisterHandler(fiber.MethodDelete, "/dev/db", auth.Public, resetDatabase)
}

func resetDatabase(c *fiber.Ctx) error {
	db := newSession(c)

	err := db.ResetDatabase()
	if err != nil {
		log.WithErr(err).Error("Failed to reset database")
		return responder.InternalServerError(c)
	}

	log.Debug("Here?")

	return responder.Ok(c)
}
