import type { User } from './User'

export interface Registration {
  // dbModel: DbModel
  // user_id: UserId
  user: User
  questions: Array<Question>
}

export interface Question {
  // dbModel: DbModel
  // registration_id: RegistrationId
  text: string
  answer: string
  render: string
}
