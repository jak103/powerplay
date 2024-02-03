package responder

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type response struct {
	StatusCode   int              `json:"status_code"`
	StatusString string           `json:"status_string"`
	Message      *string          `json:"message,omitempty"`
	Data         *json.RawMessage `json:"data,omitempty"`
}

func BadRequest(c *fiber.Ctx, message ...any) error {

	// TODO complete this
	return respond(c, fiber.StatusBadRequest, message...)
}

func NotYetImplemented(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotImplemented).SendString("Not Implemented")
}

func respond(c *fiber.Ctx, statusCode int, message ...any) error {
	msg := "TODO - Responder is not yet complete"
	res := response{
		StatusCode:   statusCode,
		StatusString: utils.StatusMessage(statusCode),
		Message:      &msg,
	}

	return c.Status(statusCode).JSON(res)
}
