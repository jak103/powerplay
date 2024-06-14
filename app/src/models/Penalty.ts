import type { GameId, PenaltyTypeId, TeamId, UserId } from './ids'
import type { DbModel, User } from './index'

export interface PenaltyType extends DbModel {
  name: string
  player_id: UserId
  player: User
  duration: number
  severity: string
}

export interface Penalty extends DbModel {
  player_id: UserId
  team_id: TeamId
  game_id: GameId
  period: number
  game_time: number
  created_by: UserId
  penalty_type: PenaltyType
  penalty_type_id: PenaltyTypeId
}

export type { PenaltyTypeId } from './ids'
