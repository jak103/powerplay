package goals

import (
	//"io"

	"github.com/gofiber/fiber/v2"
	"github.com/jak103/powerplay/internal/server/apis"
	"github.com/jak103/powerplay/internal/server/services/auth"
	"github.com/jak103/powerplay/internal/utils/responder"
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/db"


)


func init() {
	apis.RegisterHandler(fiber.MethodPost, "/goals/submit", auth.Public, postGoalsHandler)
}

func postGoalsHandler(c *fiber.Ctx) error {
	// If valid structure in post request, continue on
	goalPostRequest := new(models.Goal)
	err := c.BodyParser(goalPostRequest)
	// if err != nil {
    //     http.Error(w"Failed to read request body")
    //     return
    // }
	// Connect to database and insert goal

	if err != nil{
		return err
	}
	// if db2.connection == nil{
	// 	return db2
	// }
	// Still need to implement duration parser and goal period placment
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
	//   duration: "90m",       // replace with actual duration in a format supported by Go's time.ParseDuration
	//   period: 3,             // replace with actual period
	//   assist1_id: 101112,    // replace with actual assist1_id
	//   assist2_id: 131415,    // replace with actual assist2_id
	//   powerplay: 1,          // replace with actual powerplay status
	//   penalty: true          // replace with actual penalty status
	// };



	db := db.GetSession(c)
	record, err := db.SaveGoal(goalPostRequest)
	
	if err != nil{
		return responder.InternalServerError(c)

	}
	if record == nil {
		return responder.BadRequest(c, "Could not post goal into database")
	}
	
	// if db2.connection != nil{
	// 	result := db2.connection.Create(goalPostRequest)
	// 	if result == nil{
	// 		return result
	// 	}
	// }
	
	return responder.Ok(c)

}


// func insertGoalIntoDatabase()
// {

// }