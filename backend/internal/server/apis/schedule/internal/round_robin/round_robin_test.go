package round_robin

import (
	"reflect"
	"testing"

	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/structures"
)

// TODO finalize the round robin algorithm

func TestRoundRobin(t *testing.T) {
	t.Run("Test successful run", func(t *testing.T) {
		t.Skip("Skipping test")
	})

	t.Run("Test no leagues failure", func(t *testing.T) {
		t.Skip("Skipping test")
	})

	t.Run("Test no ice times failure", func(t *testing.T) {
		t.Skip("Skipping test")
	})

	t.Run("Test no games per team failure", func(t *testing.T) {
		t.Skip("Skipping test")
	})

	t.Run("Test generate games error", func(t *testing.T) {
		t.Skip("Skipping test")
	})

	t.Run("Test assign times error", func(t *testing.T) {
		t.Skip("Skipping test")
	})
}

func TestOptimizeSchedule(t *testing.T) {

	t.Run("Test successful run", func(t *testing.T) {
		t.Skip("Skipping test")
	})

	t.Run("Test no games failure", func(t *testing.T) {
		t.Skip("Skipping test")
	})
}

func TestGetBalanceCount(t *testing.T) {
	t.Skip("Skipping test")
}

func TestGenerateGames(t *testing.T) {
	t.Skip("Skipping test")
}

func TestAssignTimes(t *testing.T) {
	t.Skip("Skipping test")
}

func TestRotateTeams(t *testing.T) {
	t.Skip("Skipping test")
}

func TestNewGame(t *testing.T) {
	t.Skip("Skipping test")
}

func TestNewGames(t *testing.T) {
	t.Skip("Skipping test")
}

func TestReorderLeagues(t *testing.T) {

	dummyRound := []structures.Round{{Games: make([]models.Game, 0)}}

	tests := []struct {
		name     string
		roundMap map[string][]structures.Round
		want     []string
	}{
		{
			name: "simple case",
			roundMap: map[string][]structures.Round{
				"A": dummyRound,
				"B": dummyRound,
				"C": dummyRound,
				"D": dummyRound,
			},
			want: []string{"A", "C", "B", "D"},
		},
		{
			name: "odd number",
			roundMap: map[string][]structures.Round{
				"V": dummyRound,
				"W": dummyRound,
				"X": dummyRound,
				"Y": dummyRound,
				"Z": dummyRound,
			},
			want: []string{"Y", "W", "Z", "X", "V"},
		},
		{
			name: "single element",
			roundMap: map[string][]structures.Round{
				"Single": dummyRound,
			},
			want: []string{"Single"},
		},
		{
			name:     "empty map",
			roundMap: map[string][]structures.Round{},
			want:     []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reorderLeagues(tt.roundMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reorderLeagues() = %v, want %v", got, tt.want)
			}
		})
	}
}
