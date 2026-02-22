package validators

import (
	"github.com/jak103/powerplay/internal/models"
	"reflect"
	"testing"
)

type TestLeague struct {
	models.DbModel
	CorrelationId string `json:"correlation_id"`
	SeasonID      uint   `json:"season_id"`
	Name          string `json:"name"`
}

func TestIsValidSortField(t *testing.T) {
	tests := []struct {
		field       string
		expectValid bool
	}{
		{"id", true},
		{"created_at", true},
		{"updated_at", true},
		{"deleted_at", true},
		{"correlation_id", true},
		{"season_id", true},
		{"name", true},
		{"invalid_field", false},
	}

	modelType := reflect.TypeOf(TestLeague{})

	for _, test := range tests {
		t.Run(test.field, func(t *testing.T) {
			valid := IsValidSortField(test.field, modelType)
			if valid != test.expectValid {
				t.Errorf("Expected IsValidSortField(%q) to be %v, got %v", test.field, test.expectValid, valid)
			}
		})
	}
}
