package models

type User struct {
	dbModel
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Email      string  `json:"email"`
	Password   string  `json:"-"` // Password should never leave the backend
	Phone      string  `json:"phone"`
	Role       Role    `json:"role"`
	SkillLevel int     `json:"skill_level"`
	Teams      []*Team `json:"teams"`
}
