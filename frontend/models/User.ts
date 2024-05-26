import type { Role } from './types/role'
import type { Roster } from './Roster'

export interface User {
  // dbModel: DbModel
  first_name: string
  last_name: string
  email: string
  phone: string
  roles: Array<Role>
  skill_level: number
  rosters?: Array<Roster>
  staffs?: Array<Roster>
}
