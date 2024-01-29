package models

import "time"

// TODO stats needs a lot more work. Need to think better about the data model

type GameStats struct {
	Score       int       `json:"score"`
	ShotsOnGoal int       `json:"shots_on_goal"`
	Penalties   []Penalty `json:"penalties"`
}

type TeamStats struct {
	GamesStats []GameStats
}

type PlayerStats struct {
	Goals int `goals`
}

type Penalty struct {
	Type     string        `json:"type"`
	Player   User          `json:"player"`
	Duration time.Duration `json:"duration"`
}
