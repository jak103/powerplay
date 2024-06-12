package ref

// import (
// 	"fmt"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/jak103/powerplay/internal/server/apis"
// 	"github.com/jak103/powerplay/internal/server/services/auth"
// )

// func init() {
// 	apis.RegisterHandler(fiber.MethodPost, "/rsvp/ref", auth.Authenticated, handleRef)
// }

type RefScheduleRow struct {
	Start           string `csv:"Start Date and Time"`
	DurationHours   string `csv:"Duration Hours"`
	DurationMinutes string `csv:"Duration Minutes"`
	Location        string `csv:"Arena/Rink"`
	Level           string `csv:"Game Level"`
	Home            string `csv:"Home Team"`
	Away            string `csv:"Away Team"`
}
