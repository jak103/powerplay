import type { TimeString } from './types/timeString'
import type { DbModel, Registration, Game, League } from './index'

export interface Season extends DbModel {
  name: string
  start: TimeString
  end: TimeString
  registrations: Array<Registration>
  schedule: Array<Game>
  leagues: Array<League>
}

export type { SeasonId } from './ids'
