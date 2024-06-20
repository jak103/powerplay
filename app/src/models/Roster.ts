import type { DbModel, User } from './index'
import type { UserId } from './ids'

export interface Roster extends DbModel {
  players?: Array<User>
  captain: User
  captain_id: UserId
}

export type { RosterId } from './ids'
