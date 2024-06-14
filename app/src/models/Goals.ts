import type { DbModel, Game, Team } from './index'
import type { UserId, GameId, TeamId } from './ids'

export interface Goals extends DbModel {
  user_id: UserId
  game_id: GameId
  game: Game
  team_id: TeamId
  team: Team
  duration: number
  period: number
  assist1_id: UserId
  assist2_id: UserId
  playerdifferential: number
  ispenaltyshot: boolean
}
