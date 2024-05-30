package round_robin

import (
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/structures"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRoundRobin(t *testing.T) {
	t.Run("Test RoundRobin", func(t *testing.T) {
		var leagues []structures.League
		var iceTimes []string
		var numberOfGamesPerTeam int
		leagues = append(leagues, structures.League{Name: "Test", Teams: []structures.Team{{Name: "Team1", Id: "1"}, {Name: "Team2", Id: "2"}}})
		iceTimes = append(iceTimes, "5/20/24 20:30")
		numberOfGamesPerTeam = 1
		games, err := RoundRobin(leagues, iceTimes, numberOfGamesPerTeam)
		assert.Nil(t, err)
		assert.NotNil(t, games)
	})
}
