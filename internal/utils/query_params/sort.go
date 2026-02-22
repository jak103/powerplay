package query_params

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/utils/validators"
	"reflect"
	"strings"
)

func GetSortParams(c *fiber.Ctx, modelType reflect.Type) (string, string, error) {
	sortField := c.Query("sort_field", "ID")
	sortOrder := strings.ToUpper(c.Query("sort_order", "ASC"))

	if sortOrder != "ASC" && sortOrder != "DESC" {
		return "", "", fiber.NewError(fiber.StatusBadRequest, "Invalid sort_order parameter")
	}

	if !validators.IsValidSortField(sortField, modelType) {
		return "", "", fiber.NewError(fiber.StatusBadRequest, "Invalid sort_field parameter")
	}

	return sortField, sortOrder, nil
}
