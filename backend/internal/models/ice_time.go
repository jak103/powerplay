package models

type IceTime struct {
	DbModel
	IceTimeID int    `json:"ice_time_id"`
	StartTime string `json:"start_time"`
}
