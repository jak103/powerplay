import type { UserId } from './types/ids'
import type { User } from './User'
import type { DurationNs } from './types/durationNs'

export interface PenaltyType {
  // dbModel: DbModel
  name: string
  player_id: UserId
  player: User
  duration: DurationNs
  severity: string
}
