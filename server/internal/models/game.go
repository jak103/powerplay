package models

import "time"

type Game struct {
	dbModel
	Home           Team      `json:"home"`
	Away           Team      `json:"away"`
	HomeLockerRoom string    `json:"home_locker_room"`
	AwayLockerRoom string    `json:"away_locker_room"`
	Start          time.Time `json:"start"`
	End            time.Time `json:"end"`
	Location       Venue     `json:"location"`
	Staff          Staff     `json:"staff"`
	HomeScore      int       `json:"home_score"`
	AwayScore      int       `json:"away_score"`
}
