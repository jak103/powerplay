package models

import "time"

type Season struct {
	dbModel
	Name          string         `json:"name"`
	Leagues       []League       `json:"leagues"`
	Start         time.Time      `json:"start"`
	End           time.Time      `json:"end"`
	Registrations []Registration `json:"registrations"`
}
