package models

import "time"

type ShotOnGoal struct{
	DbModel
	GameId uint 'json:"game_id" gorm:"not_null"'
	TeamId uint 'json:"team_id" gorm:"not_null"'
	ShotTime uint 'json:"shot_time"'
	Scorekeeper uint 'json:"scorekeeper gorm:"not_null"'
}