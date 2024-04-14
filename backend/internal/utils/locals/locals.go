package locals

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/utils/constants"
	"github.com/jak103/powerplay/internal/utils/log"
)

func SetLogger(c *fiber.Ctx, l log.Logger) {
	c.Locals("logger", l)
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

func SetRequestingUser(c *fiber.Ctx, user models.User) {
	c.Locals(constants.RequestingUserLocal, user)
}

func RequestingUser(c *fiber.Ctx) *models.User {
	var user *models.User
	if c != nil {
		if r := c.Locals(constants.RequestingUserLocal); r != nil {
			temp := r.(models.User)
			user = &temp
		}
	}
	return user
}
