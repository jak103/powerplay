package models

import "github.com/jak103/powerplay/internal/server/services/auth"

type KeyRecord struct {
	DbModel
	UserId uint        `json:"user_id"`
	Roles  []auth.Role `json:"roles" gorm:"type:text[]"`
}
