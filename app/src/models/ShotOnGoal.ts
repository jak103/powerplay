import type { DbModel } from "./index"
import type { GameId, TeamRecordId, UserId } from "./ids"

export interface ShotOnGoal extends DbModel {
  game_id: GameId
  team_id: TeamRecordId
  shot_time: number
  scorekeeper: UserId
}
