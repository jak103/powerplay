package models

import (
	"time"

	"github.com/jak103/powerplay/internal/server/services/auth"
)

type User struct {
	DbModel
	FirstName    string      `json:"first_name"`
	LastName     string      `json:"last_name"`
	Email        string      `json:"email"`
	Password     string      `json:"-"` // Password should never leave the backend
	Phone        string      `json:"phone"`
	Role         []auth.Role `json:"roles" gorm:"type:text[]"`
	SkillLevel   int         `json:"skill_level"`
	CurrentTeams []Team      `json:"current_teams" gorm:"many2many:users_teams"`
	DateOfBirth  time.Time   `json:"dob"`
}
