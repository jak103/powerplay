import type { DbModel, Team } from './index'
import type { LeagueGuid, SeasonId } from './ids'

export interface League extends DbModel {
  correlation_id: LeagueGuid
  season_id: SeasonId
  name: string
  teams: Array<Team>
}

export type { LeagueId, LeagueGuid } from './ids'
