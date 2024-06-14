import { RosterId, SeasonId, TeamId, UserId, VenueId } from './ids'
import type { DbModel, Roster, Team, User, Venue } from './index'
import { TimeString } from './types/timeString'


export enum GameStatus {
  SCHEDULED = 'Scheduled',
  IN_PROGRESS = 'In Progress',
  FINAL = 'Final'
}

export interface Game extends DbModel {
  season_id: SeasonId
  start: TimeString
  venue: Venue
  venue_id: VenueId
  status: GameStatus
  
  // TODO: This might change to home: { ... } and away: { ... }
  home_team: Team
  home_team_id: TeamId
  home_team_roster: Roster
  home_team_roster_id: RosterId
  home_team_locker_room: string
  home_team_shots_on_goal: number
  home_team_score: number
  
  away_team: Team
  away_team_id: TeamId
  away_team_roster: Roster
  away_team_roster_id: RosterId
  away_team_locker_room: string
  away_team_shots_on_goal: number
  away_team_score: number
  
  score_keeper: User
  score_keeper_id: UserId
  primary_referee: User
  primary_referee_id: UserId
  secondary_referee: User
  secondary_referee_id: UserId
}

export type { GameId } from './ids'
