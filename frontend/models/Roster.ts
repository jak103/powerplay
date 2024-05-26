import type { Team } from './Team'
import type { Season } from './Season'
import type { User } from './User'

export interface Roster {
  // dbModel: DbModel
  // season_id: SeasonId
  team: Team
  // TeamID: TeamId
  season: Season
  players: Array<User>
  staff: Array<User>
}
