import type { DbModel, Team } from './index'
import type { Role } from './KeyRecord'
import type { TimeString } from './types/timeString'

export interface User extends DbModel {
  first_name: string
  last_name: string
  email: string
  phone: string
  roles: Array<Role>
  skill_level: number
  current_teams: Array<Team>
  /** dob === Date of Birth */
  dob: TimeString
}

export type { UserId } from './ids'
