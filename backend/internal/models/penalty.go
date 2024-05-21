package models

import "time"

type PenaltyType struct {
	DbModel
	Name     string        `json:"name"`
	Duration time.Duration `json:"duration"`
	Severity string        `json:"severity"`
}
