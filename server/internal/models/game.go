package models

import "time"

type Game struct {
	dbModel
	Home     Team      `json:"home"`
	Away     Team      `json:"away"`
	Start    time.Time `json:"start"`
	End      time.Time `json:"end"`
	Location Venue     `json:"location"`
}
