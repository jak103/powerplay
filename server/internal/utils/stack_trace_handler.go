package utils

import (
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/leaguemanager/internal/utils/locals"
)

func StackTraceHandler(c *fiber.Ctx, e interface{}) {
	log := locals.Logger(c)
	log.Error("panic: %v\n%s\n", e, debug.Stack())
}
