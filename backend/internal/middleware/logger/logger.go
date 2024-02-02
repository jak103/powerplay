package logger

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/utils/locals"
	"github.com/jak103/powerplay/internal/utils/log"
)

func New() fiber.Handler {
	return middleware
}

func middleware(c *fiber.Ctx) error {
	reqid := locals.RequestId(c)
	logger := log.WithRequestId(reqid)

	locals.SetLogger(c, logger)

	logger.Info("Starting %s %s", c.Method(), c.OriginalURL())
	start := time.Now()
	err := c.Next()
	total := time.Since(start)

	log.Info("Finished %s %s -- [%v] %v", c.Method(), c.Path(), c.Response().StatusCode(), total)
	return err
}
