import type { TimeString } from './types/timeString'
import type { Game } from './Game'
import type { Roster } from './Roster'

export interface Season {
  // dbModel: DbModel
  name: string
  start: TimeString
  end: TimeString
  schedule: Array<Game>
  // registrations: Array<Registration>
  rosters: Array<Roster>
}
