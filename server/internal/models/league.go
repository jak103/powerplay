package models

type Division struct {
	Name  string `json:"name"`
	Teams []Team `json:"teams"`
}
