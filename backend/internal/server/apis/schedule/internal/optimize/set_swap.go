package optimize

import (
	"sort"

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
	teams := make([]string, 0)
	for key := range teamStats {
		teams = append(teams, key)
	}
	sort.Strings(teams)

	for i := 0; i < 20; i++ { // FIXME hardcoded 20 for number of "league sets"
		relevant_games := games[i : i+6]
		worst_imbalance := -1
		worst_index := -1
		var worst_stats structures.TeamStats
		for j := 0; j < 6; j++ {
			team0 := relevant_games[j].HomeTeam.Name
			team1 := relevant_games[j].AwayTeam.Name
			team0stats := teamStats[team0]
			team1stats := teamStats[team1]

			team0delta := team0stats.EarlyGames - team0stats.LateGames
			if team0delta < 0 {
				team0delta *= -1
			}
			if team0delta > worst_imbalance {
				worst_stats = team0stats
				worst_imbalance = team0delta
				worst_index = i*6 + j
			}

			team1delta := team1stats.EarlyGames - team1stats.LateGames
			if team1delta < 0 {
				team1delta *= -1
			}
			if team1delta > worst_imbalance {
				worst_stats = team1stats
				worst_imbalance = team1delta
				worst_index = i*6 + j
			}
		}

		// if all teams are balanced already don't do anything
		if worst_imbalance <= 2 {
			continue
		}

		// now that we've found the team with the worst stats this iteration, we try to find the team
		// with the most opposite imbalance to swap with
		more_early := worst_stats.EarlyGames > worst_stats.LateGames
		opposite_imbalance := -1
		opposite_index := -1
		for j := 0; j < 6; j++ {
			team0 := relevant_games[j].HomeTeam.Name
			team1 := relevant_games[j].AwayTeam.Name
			team0stats := teamStats[team0]
			team1stats := teamStats[team1]

			if more_early {
				team0delta := team0stats.LateGames - team0stats.EarlyGames
				if team0delta > opposite_imbalance {
					opposite_imbalance = team0delta
					opposite_index = i*6 + j
				}

				team1delta := team1stats.LateGames - team1stats.EarlyGames
				if team1delta > opposite_imbalance {
					opposite_imbalance = team1delta
					opposite_index = i*6 + j
				}
			} else {
				team0delta := team0stats.EarlyGames - team0stats.LateGames
				if team0delta > opposite_imbalance {
					opposite_imbalance = team0delta
					opposite_index = i*6 + j
				}

				team1delta := team1stats.EarlyGames - team1stats.LateGames
				if team1delta > opposite_imbalance {
					opposite_imbalance = team1delta
					opposite_index = i*6 + j
				}
			}
		}

		updateStats(teamStats, games, worst_index, opposite_index)
		swapGames(games, worst_index, opposite_index)
	}
}
