package models

import "time"

type Season struct {
	dbModel
	Name      string     `json:"name"`
	Divisions []Division `json:"divisions"`
	Start     time.Time  `json:"start"`
	End       time.Time  `json:"end"`
}
