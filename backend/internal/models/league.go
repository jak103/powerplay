package models

type League struct {
	DbModel
	Season string //`json:"season"`
	Name   string //`json:"name"`
	Teams  []Team //`json:"teams"`//
}
