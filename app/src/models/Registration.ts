import type { DbModel, User } from './index'
import type { RegistrationId, UserId } from './ids'

export interface Registration extends DbModel {
  user_id: UserId // TODO: Why does this not have a json key tag in the go model?
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
