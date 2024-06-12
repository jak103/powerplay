package structures

import (
	"github.com/jak103/powerplay/internal/models"
	"time"
)

type Season struct {
	LeagueRounds map[string][]Round
}

type Round struct {
	Games []models.Game
}

type TeamStats struct {
	Name                    string               `json:"name"`
	League                  string               `json:"league"`
	EarlyGames              int                  `json:"earlyGame"`
	LateGames               int                  `json:"lateGame"`
	DaysOfTheWeek           map[time.Weekday]int `json:"daysOfTheWeek"`
	DaysBetweenGames        []int                `json:"daysBetweenGames"`
	AverageDaysBetweenGames float32              `json:"averageDaysBetweenGames"`
	Balanced                bool                 `json:"balanced"`
}

type SeasonStats struct {
	EarlyGames int
	LateGames  int
	TotalGames int
	Score      float64
}

func (ss SeasonStats) EarlyPercentage() float32 {
	return float32(ss.EarlyGames) / float32(ss.TotalGames)
}
