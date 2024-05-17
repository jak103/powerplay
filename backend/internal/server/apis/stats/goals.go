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
	//goalPostRequest := new(models.Goal)
	//err := c.BodyParser(goalPostRequest)
	// if err != nil {
    //     http.Error(w"Failed to read request body")
    //     return
    // }
	// Connect to database and insert goal

	// if err != nil{
	// 	return err
	// }
	// if db2.connection == nil{
	// 	return db2
	// }


	db := db.GetSession(c)
	record, err := db.SaveGoal(new(models.Goal))
	
	if err != nil{
		return responder.InternalServerError(c)

	}
	if &record == nil {
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