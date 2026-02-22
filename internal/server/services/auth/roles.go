package auth

import (
	"database/sql/driver"
	"errors"
)

type Role string

const (
	None        Role = "none"
	Player      Role = "player"
	Captain     Role = "captain"
	Referee     Role = "referee"
	ScoreKeeper Role = "scorekeeper"
	Manager     Role = "manager"
)

func (r Role) Value() (driver.Value, error) {
	return string(r), nil
}

func (r *Role) Scan(value interface{}) error {
	if value == nil {
		*r = ""
		return nil
	}
	val, ok := value.([]byte)
	if !ok {
		return errors.New("invalid type for Role")
	}
	*r = Role(val)
	return nil
}

var (
	Public        []Role = []Role{None}
	Authenticated []Role = []Role{Manager, Referee, ScoreKeeper, Captain, Player}
	Staff         []Role = []Role{Manager, Referee, ScoreKeeper}
	ManagerOnly   []Role = []Role{Manager}
)

func HasCorrectRole(usersRoles []Role, roles []Role) bool {
	for _, usersRole := range usersRoles {
		for _, neededRole := range roles {
			if neededRole == None || usersRole == neededRole {
				return true
			}
		}
	}

	return false
}
