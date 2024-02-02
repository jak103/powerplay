package models

type Venue struct {
	dbModel
	Name        string   `json:"name"`
	Address     string   `json:"address"`
	LockerRooms []string `json:"locker_rooms"`
}
