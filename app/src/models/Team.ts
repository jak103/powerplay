import type { DbModel, League, Roster } from './index'
import type { LeagueId, LogoId, RosterId, TeamGuid } from './ids'

export interface Team extends DbModel {
  correlation_id: TeamGuid
  name: string
  logo_id: LogoId
  color: string
  league_id: LeagueId
  league: League
  roster?: Roster
  roster_id?: RosterId
  wins: number
  losses: number
}

export type { TeamId, TeamGuid } from './ids'
