import type { DbModel, LeagueRecord, Roster } from './index'
import type { LeagueRecordId, LogoId, RosterId, TeamGuid } from './ids'

export interface TeamRecord extends DbModel {
  team_guid: TeamGuid
  name: string
  logo_id: LogoId
  color: string
  league_record_id: LeagueRecordId
  league_record: LeagueRecord
  roster: Roster
  roster_id: RosterId
  wins: number
  losses: number
}

export type { TeamRecordId, TeamGuid } from './ids'
