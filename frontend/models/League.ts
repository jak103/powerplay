import type { Team } from './Team'

export interface League {
  // dbModel: DbModel
  name: string
  teams: Array<Team>
}
