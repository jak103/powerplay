package goals

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/responder"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/db"
	"github.com/jak103/powerplay/internal/utils/locals"
)


func init() {
	apis.RegisterHandler(fiber.MethodPost, "/goals/submit", auth.Public, postGoalsHandler)
}

func postGoalsHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)
	log.Info("Handling creating new goal")
	log.Debug("body: %q", c.Request().Body())
	goalPostRequest := new(models.Goal)
	err := c.BodyParser(goalPostRequest)
	
	// If valid structure in post request, continue on
	if err != nil{
		log.WithErr(err).Error("Failed to parse Goal POST request.")
		return err
	}

	// Connect to database and insert goal
	db := db.GetSession(c)
	record, err := db.SaveGoal(goalPostRequest)
	
	if err != nil{
		log.WithErr(err).Alert("Failed to parse goal request payload")
		return responder.InternalServerError(c)
	}

	if record == nil {
		return responder.BadRequest(c, "Could not post goal into database")
	}
	
	return responder.Ok(c)

}

// Example api request: fetch('http://localhost:9001/api/v1/goals/submit', {
	//     method: 'POST',
	//     headers: {
	//         'Content-Type': 'application/json'
	//     },
	//     body: JSON.stringify(data)
	// })
	// .then(response => response.json())
	// .then(data => console.log(data))
	// .catch(error => console.error('Error:', error));

	// const data = {
	//   user_id: 123,          // replace with actual user_id
	//   game_id: 456,          // replace with actual game_id
	//   team_id: 789,          // replace with actual team_id
	//   duration: 15,       // replace with actual duration in a format supported by Go's time.ParseDuration
	//   period: 3,             // replace with actual period
	//   assist1_id: 101112,    // replace with actual assist1_id
	//   assist2_id: 131415,    // replace with actual assist2_id
	//   powerplay: 1,          // replace with actual powerplay status
	//   penalty: true          // replace with actual penalty status
	// };

