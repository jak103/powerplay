import type { TimeString } from './types/timeString'

/**
 * This is not necessarily the base model, but it *is* used in several other models
 */
export interface DbModel {
  id: number
  created_at: TimeString
  updated_at: TimeString
  deleted_at?: TimeString
}
