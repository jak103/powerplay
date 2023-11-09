package models

type Team struct {
	dbModel
	Name string `json:"name"`
}
