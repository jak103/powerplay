package query_params

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetPaginationParams(c *fiber.Ctx) (int, int, bool, error) {
	offsetParam := c.Query("offset", "0")
	limitParam := c.Query("limit", "10")
	fetchAll := c.Query("fetch_all", "false")

	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		return 0, 0, false, fiber.NewError(fiber.StatusBadRequest, "Invalid offset parameter")
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil {
		return 0, 0, false, fiber.NewError(fiber.StatusBadRequest, "Invalid limit parameter")
	}

	fetchAllBool, err := strconv.ParseBool(fetchAll)
	if err != nil {
		return 0, 0, false, fiber.NewError(fiber.StatusBadRequest, "Invalid fetch_all parameter")
	}

	return offset, limit, fetchAllBool, nil
}
