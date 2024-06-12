import type { DbModel } from './index'
import type { UserId } from './ids'

export enum Role {
  None = 'none',
  Player = 'player',
  Captain = 'captain',
  Referee = 'referee',
  ScoreKeeper = 'scorekeeper',
  Manager = 'manager'
}

export interface KeyRecord extends DbModel {
  user_id: UserId
  roles: Array<Role>
}
