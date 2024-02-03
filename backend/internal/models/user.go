package models

type User struct {
	DbModel
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Email      string    `json:"email"`
	Password   string    `json:"-"` // Password should never leave the backend
	Phone      string    `json:"phone"`
	Role       Role      `json:"role"`
	SkillLevel int       `json:"skill_level"`
	Rosters    []*Roster `json:"rosters,omitempty" gorm:"many2many:player_rosters"`
	Staffs     []*Roster `json:"staffs,omitempty" gorm:"many2many:staff_rosters"`
}
