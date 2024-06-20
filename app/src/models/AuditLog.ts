import { AuditLogId } from './ids'
import type { User } from './index'

/**
 * @unused We could use this if we need an admin frontend.
 */
export interface AuditLog {
  user: User
  action: string
  table: string
  id: AuditLogId
}

export type { AuditLogId } from './ids'
