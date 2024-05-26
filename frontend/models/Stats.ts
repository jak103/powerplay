import type { DurationNs } from './types/durationNs'
import type { User } from './User'

/**
 * Game stats are subject to change.
 * 
 * Any properties can be assigned to this type,
 * but to use them, we have to ignore TS warnings.
 */
export type GameStats = {
  score: number
  shots_on_goal: number
  penalities: Array<PenaltyStat>
} & Record<string, unknown>

/**
 * Team stats are subject to change.
 * 
 * Any properties can be assigned to this type,
 * but to use them, we have to ignore TS warnings.
 */
export type TeamStats = {
  // GamesStats: Array<GameStats>
} & Record<string, unknown>

/**
 * Player stats are subject to change.
 * 
 * Any properties can be assigned to this type,
 * but to use them, we have to ignore TS warnings.
 */
export type PlayerStats = {
  goals: number
} & Record<string, unknown>

/**
 * Penalty stats are subject to change.
 * 
 * Any properties can be assigned to this type,
 * but to use them, we have to ignore TS warnings.
 */
export type PenaltyStat = {
  type: string
  player: User
  duration: DurationNs
} & Record<string, unknown>
