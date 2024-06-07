import type { DbModel, Game, TeamRecord } from './index'
import type { UserId, GameId, TeamRecordId } from './ids'

export interface Goals extends DbModel {
  user_id: UserId
  game_id: GameId
  game: Game
  team_record_id: TeamRecordId
  team_record: TeamRecord
  duration: number
  period: number
  assist1_id: UserId
  assist2_id: UserId
  player_differential: number
  is_penalty_shot: boolean
}
