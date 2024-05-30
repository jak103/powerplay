import type { DbModel } from "./DbModel"
import type { User } from "./index"

/**
 * @unused We could use this if we need an admin frontend.
 */
export interface AuditLog extends DbModel {
  user: User
  action: string
  table: string
}

export type { AuditLogId } from './ids'
