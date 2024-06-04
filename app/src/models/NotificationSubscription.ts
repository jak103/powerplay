import type { DbModel } from './index'

export enum NotifTopic {
  RSVP = 'new_rsvp',
  CHAT = 'new_chat',
  GAME_UPDATE = 'game_update',
  EVENT_UPDATE = 'event_update'
}

export interface NotificationSubscription extends DbModel {
  topics: Array<NotifTopic>
  // TODO: Are we supposed to be getting this webpush stuff?
  endpoint: string
  auth: string
  p265dh: string
}
