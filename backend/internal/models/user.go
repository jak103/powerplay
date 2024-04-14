package models

import (
	"time"

	"github.com/jak103/powerplay/internal/server/services/auth"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Experience struct {
	SkillLevel   int `json:"skill_level"`
	CoachedTime  int `json:"coached_time"`
	TimeInLeague int `json:"time_in_league"`
	// TODO Update this when we need to calculate skill levels
}

type User struct {
	DbModel
	FirstName      string      `json:"first_name"`
	LastName       string      `json:"last_name"`
	Email          string      `json:"email"`
	Phone          string      `json:"phone"`
	Birthdate      time.Time   `json:"birth_date"`
	HashedPassword string      `json:"-"` // Password never leaves the backend
	Salt           string      `json:"-"` // Salt never leaves the backend
	Roles          []auth.Role `json:"roles" gorm:"type:string"`
	Experience     Experience  `json:"experience" gorm:"embedded"`
	Rosters        []*Roster   `json:"rosters,omitempty" gorm:"many2many:player_rosters"`
	Staffs         []*Roster   `json:"staffs,omitempty" gorm:"many2many:staff_rosters"`
	Verified       bool        `json:"verified"`
}
