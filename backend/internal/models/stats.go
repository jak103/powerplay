package models

import "time"

// TODO stats needs a lot more work. Need to think better about the data model

type SeasonStats struct {
	EarlyGames int
	LateGames  int
	TotalGames int
}

func (ss SeasonStats) EarlyPercentage() float32 {
	return float32(ss.EarlyGames) / float32(ss.TotalGames)
}

type GameStats struct {
	Score       int       `json:"score"`
	ShotsOnGoal int       `json:"shots_on_goal"`
	Penalties   []Penalty `json:"penalties"`
}

type TeamStats struct {
	GamesStats              []GameStats
	Name                    string
	League                  string
	EarlyGames              int
	LateGames               int
	DaysOfTheWeek           map[time.Weekday]int
	DaysBetweenGames        []int
	AverageDaysBetweenGames float32
	Games                   []Game
	Balanced                bool
}

type PlayerStats struct {
	Goals int `goals`
}

type Penalty struct {
	Type     string        `json:"type"`
	Player   User          `json:"player"`
	Duration time.Duration `json:"duration"`
}
