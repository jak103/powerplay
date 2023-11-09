package models

type Player struct {
	dbModel
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
