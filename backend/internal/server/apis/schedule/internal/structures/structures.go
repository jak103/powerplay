package structures

import (
	"github.com/jak103/powerplay/internal/models"
	"time"
)

// TODO need to move these into the structures package in the internal directory
// TODO need to update the structures I created based off of these

type Season struct {
	LeagueRounds map[string][]Round
}

type Round struct {
	Games []Game
}

// https://help.sportsengine.com/en/articles/6380725-schedule-upload-tutorial
// https://intercom.help/SportsEngine/en/articles/6310600-schedule-upload-fields-reference-guide
type Game struct {
	Optimized bool `json:"optimized"`
	models.Game
}

//func (g Game) String() string {
//	return fmt.Sprintf("%s %s v %s", g.Start.Format("01-02-2006 03:04"), g.Team1Name, g.Team2Name)
//	// DateOnly   = "2006-01-02"
//	// TimeOnly   = "15:04:05"
//}

type LeagueIceTimes struct {
	Name  string   `yaml:"league"`
	Times []string `yaml:"times"`
	Used  []bool   `yaml:"-"`
}

type OptimizeGameWrapper struct {
	models.Game
}

type TeamStats struct {
	Name                    string               `json:"name"`
	League                  string               `json:"league"`
	EarlyGames              int                  `json:"earlyGame"`
	LateGames               int                  `json:"lateGame"`
	DaysOfTheWeek           map[time.Weekday]int `json:"daysOfTheWeek"`
	DaysBetweenGames        []int                `json:"daysBetweenGames"`
	AverageDaysBetweenGames float32              `json:"averageDaysBetweenGames"`
	Games                   []Game               `json:"games"`
	Balanced                bool                 `json:"balanced"`
}

type SeasonStats struct {
	EarlyGames int
	LateGames  int
	TotalGames int
}

func (ss SeasonStats) EarlyPercentage() float32 {
	return float32(ss.EarlyGames) / float32(ss.TotalGames)
}
