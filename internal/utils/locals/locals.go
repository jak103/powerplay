package locals

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/utils/constants"
	"github.com/jak103/powerplay/internal/utils/log"
)

const (
	LOGGER = "logger"
)

func SetLogger(c *fiber.Ctx, l log.Logger) {
	c.Locals(LOGGER, l)
}

func Logger(c *fiber.Ctx) log.Logger {
	var logger log.Logger

	if c == nil {
		logger = log.TheLogger
	} else if l := c.Locals("logger"); l != nil {
		logger = l.(log.Logger)
	}

	return logger
}

func RequestId(c *fiber.Ctx) string {
	reqid := "missing_request_id"
	// TODO Do I really need a nil check here?
	if c != nil {
		if r := c.Locals(constants.RequestIdLocal); r != nil {
			reqid = r.(string)
		}
	}
	return reqid
}

func SetKeyRecord(c *fiber.Ctx, record models.KeyRecord) {
	c.Locals(constants.KeyRecordLocal, record)
}

func KeyRecord(c *fiber.Ctx) *models.KeyRecord {
	var record *models.KeyRecord
	if c != nil {
		if r := c.Locals(constants.KeyRecordLocal); r != nil {
			temp := r.(models.KeyRecord)
			record = &temp
		}
	}
	return record
}
