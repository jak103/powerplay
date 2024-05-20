package schedule

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
)

func Init() {
	apis.RegisterHandler(fiber.MethodPost, "/schedule/generate", auth.Authenticated, HandleGenerate)
	apis.RegisterHandler(fiber.MethodPost, "/schedule/analysis", auth.Authenticated, handleAnalysis)
	apis.RegisterHandler(fiber.MethodPost, "/schedule/ref", auth.Authenticated, handleRef)
	apis.RegisterHandler(fiber.MethodPost, "/rsvp", auth.Authenticated, handleRsvp)
	apis.RegisterHandler(fiber.MethodPost, "/schedule/update_ids", auth.Authenticated, handleUpdateIds)
}
