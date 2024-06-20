import type { DbModel } from './index'
import type { GameId, TeamId, UserId } from './ids'

export interface ShotOnGoal extends DbModel {
  game_id: GameId
  team_id: TeamId
  shot_time: number
  scorekeeper: UserId
}
