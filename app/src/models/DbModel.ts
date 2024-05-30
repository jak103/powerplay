import type { TimeString } from './types/timeString'

/**
 * A base model. Most models extend this one.
 */
export interface DbModel {
  id: number
  created_at: TimeString
  updated_at: TimeString
  deleted_at?: TimeString
}
