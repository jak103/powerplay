import type { NotifTopic } from "./types/notifTopic";

export interface NotificationSubscription {
  // dbModel: DbModel
  topics: Array<NotifTopic>
}
