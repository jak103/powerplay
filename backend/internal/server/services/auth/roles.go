package auth

type Role string

const (
	// Roles
	None        Role = "none"
	Player      Role = "player"
	Captain     Role = "captain"
	Referee     Role = "referee"
	ScoreKeeper Role = "scorekeeper"
	Manager     Role = "manager"
)

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
