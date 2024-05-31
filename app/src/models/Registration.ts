import type { DbModel, User } from './index'
import type { RegistrationId, SeasonId, UserId } from './ids'

export interface Registration extends DbModel {
  user_id: UserId
  season_id: SeasonId
  user: User
  questions: Array<Question>
}

export interface Question extends DbModel {
  registration_id: RegistrationId
  text: string
  answer: string
  render: string
}

export type { RegistrationId } from './ids'
