import type { Team } from './Team'
import type { TimeString } from './types/timeString'
import type { Venue } from './Venue'

export interface Game {
  // dbModel: DbModel
  // season_id: SeasonId
  teams: Array<Team>
  home_locker_room: string
  away_locker_room: string
  start: TimeString
  end: TimeString
  venue: Venue
  // venue_id: VenueId

  // manager_on_call: User
  // // manager_on_call_id: UserId
  // score_keeper: User
  // primary_referee: User
  // secondary_referee: User

  // May be replaced with stats models
  // home_score: number
  // away_score: number
}
