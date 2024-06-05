import type { DbModel, TeamRecord } from './index'
import type { LeagueGuid, SeasonId } from './ids'

export interface LeagueRecord extends DbModel {
  league_guid: LeagueGuid
  season_id: SeasonId
  name: string
  teams: Array<TeamRecord>
}

export type { LeagueRecordId, LeagueGuid } from './ids'
