package stats

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
)

func init() {
	apis.RegisterHandler(fiber.MethodGet, "/penaltyTypes", auth.Public, getPenaltyTypes)
	apis.RegisterHandler(fiber.MethodGet, "/penalties", auth.Public, getPenalties)
	//apis.RegisterHandler(fiber.MethodPost, "/penalty", auth.Public, createPenalty)
}

func getPenaltyTypes(c *fiber.Ctx) error {

	penaltyTypes := stubPenaltyTypes()
	jsonData, err := json.Marshal(penaltyTypes)
	if err != nil {
		return err
	}

	c.Type("json")

	// Send JSON response
	return c.Send(jsonData)
}

func stubPenaltyTypes() map[string][]string {
	return map[string][]string{"penaltyTypes": {"Boarding", "Charging", "Slashing"}}
}

func getPenalties(c *fiber.Ctx) error {
	penalties := stubPenalties()
	jsonData, err := json.Marshal(penalties)
	if err != nil {
		return err
	}

	c.Type("json")

	// Send JSON response
	return c.Send(jsonData)
}

func stubPenalties() []models.Penalty {
	var penalties []models.Penalty
	jsonString := `[{"id":1,"created_at":"2009-11-10T23:00:00Z","updated_at":"2009-11-10T23:00:00Z","player_id":55,"team_id":1,"game_id":42,"period":2,"duration":123,"created_by":5,"penalty_type_id":1},{"id":2,"created_at":"2009-11-10T23:00:00Z","updated_at":"2009-11-10T23:00:00Z","player_id":55,"team_id":1,"game_id":42,"period":3,"duration":456,"created_by":5,"penalty_type_id":1}]`

	err := json.Unmarshal([]byte(jsonString), &penalties)
	if err != nil {

		// if error is not nil
		// print error
		fmt.Println(err)
	}
	return penalties
}

// func createPenalty(c *fiber.Ctx) error {
// }
