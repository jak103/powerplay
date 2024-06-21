package auth

import (
	"database/sql/driver"
	"errors"
	"strings"
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

func HasCorrectRole(usersRoles Roles, roles Roles) bool {
	for _, usersRole := range usersRoles {
		for _, neededRole := range roles {
			if neededRole == None || usersRole == neededRole {
				return true
			}
		}
	}

	return false
}

type Roles []Role

var (
	Public        Roles = []Role{None}
	Authenticated Roles = []Role{Manager, Referee, ScoreKeeper, Captain, Player}
	Staff         Roles = []Role{Manager, Referee, ScoreKeeper}
	ManagerOnly   Roles = []Role{Manager}
)

func (rs *Roles) Scan(value interface{}) error {
	if value == nil {
		*rs = nil
		return nil
	}

	var rolesStr []string
	switch v := value.(type) {
	case []byte:
		rolesStr = strings.Split(string(v), ",")
	case string:
		rolesStr = strings.Split(v, ",")
	default:
		return errors.New("unsupported type for Roles")
	}

	roles := make(Roles, len(rolesStr))
	for i, roleStr := range rolesStr {
		roles[i] = Role(roleStr)
	}

	*rs = roles
	return nil
}

func (rs Roles) Value() (driver.Value, error) {
	roles := make([]string, len(rs))

	for i, role := range rs {
		roles[i] = string(role)
	}

	return roles, nil
}
