import type { DbModel } from "./index"

export interface Venue extends DbModel{
  name: string
  address: string
  locker_rooms: Array<string>
}

export type { VenueId } from './ids'
