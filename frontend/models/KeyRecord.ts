import type { UserId } from './types/ids'
import type { Role } from './types/role'

export interface KeyRecord {
  // dbModel: DbModel
  user_id: UserId
  roles: Array<Role>
}
