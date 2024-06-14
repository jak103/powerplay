package optimize

import (
	"github.com/jak103/powerplay/internal/models"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/analysis"
	"github.com/jak103/powerplay/internal/server/apis/schedule/internal/structures"
	"github.com/jak103/powerplay/internal/utils/log"
)

func SetOptimizeSchedule(games []models.Game) {
	if len(games) == 0 {
		log.Info("No games to optimize")
		return
	}
	log.Info("Pre-optimization analysis")
	seasonStats, teamStats := analysis.RunTimeAnalysis(games)
	pairSwapSchedule(games, seasonStats, teamStats)
	log.Info("Finished running optimizer")
}

func setSwapSchedule(games []models.Game, seasonStats structures.SeasonStats, teamStats map[string]structures.TeamStats) {
	var leagues map[string]bool // set of leagues
	for _, team := range teamStats {
		if _, ok := leagues[team.League]; !ok {
			leagues[team.League] = true
		}
	}
	teamsPerLeague := len(teamStats) / len(leagues)

	// This algorithm is based on taking sets of games and trying to swap games for teams with inverse imbalances
	for i := 0; i < len(games)/teamsPerLeague; i++ {
		relevantGames := games[i : i+teamsPerLeague]
		worstImbalance := -1
		worstIndex := -1
		var worstStats structures.TeamStats
		for j := 0; j < teamsPerLeague; j++ {
			team0 := relevantGames[j].HomeTeam.Name
			team1 := relevantGames[j].AwayTeam.Name
			team0stats := teamStats[team0]
			team1stats := teamStats[team1]

			team0delta := team0stats.EarlyGames - team0stats.LateGames
			if team0delta < 0 {
				team0delta *= -1
			}
			if team0delta > worstImbalance {
				worstStats = team0stats
				worstImbalance = team0delta
				worstIndex = i*teamsPerLeague + j
			}

			team1delta := team1stats.EarlyGames - team1stats.LateGames
			if team1delta < 0 {
				team1delta *= -1
			}
			if team1delta > worstImbalance {
				worstStats = team1stats
				worstImbalance = team1delta
				worstIndex = i*teamsPerLeague + j
			}
		}

		// if all teams are balanced already don't do anything
		if worstImbalance <= 2 {
			continue
		}

		// now that we've found the team with the worst stats this iteration, we try to find the team
		// with the most opposite imbalance to swap with
		moreEarly := worstStats.EarlyGames > worstStats.LateGames
		oppositeImbalance := -1
		oppositeIndex := -1
		for j := 0; j < teamsPerLeague; j++ {
			team0 := relevantGames[j].HomeTeam.Name
			team1 := relevantGames[j].AwayTeam.Name
			team0stats := teamStats[team0]
			team1stats := teamStats[team1]

			if moreEarly {
				team0delta := team0stats.LateGames - team0stats.EarlyGames
				if team0delta > oppositeImbalance {
					oppositeImbalance = team0delta
					oppositeIndex = i*teamsPerLeague + j
				}

				team1delta := team1stats.LateGames - team1stats.EarlyGames
				if team1delta > oppositeImbalance {
					oppositeImbalance = team1delta
					oppositeIndex = i*teamsPerLeague + j
				}
			} else {
				team0delta := team0stats.EarlyGames - team0stats.LateGames
				if team0delta > oppositeImbalance {
					oppositeImbalance = team0delta
					oppositeIndex = i*teamsPerLeague + j
				}

				team1delta := team1stats.EarlyGames - team1stats.LateGames
				if team1delta > oppositeImbalance {
					oppositeImbalance = team1delta
					oppositeIndex = i*teamsPerLeague + j
				}
			}
		}

		updateStats(teamStats, games, worstIndex, oppositeIndex)
		swapGames(games, worstIndex, oppositeIndex)
	}
}
