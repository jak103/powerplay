import type { Branded } from 'src/ts_helpers/branding'

/**
 * `TimeString` is the same as a string at runtime.
 * @see {@link Branded}
 */
export type TimeString = Branded<string, 'time_string'>

export const TimeString = {
  /**
   * Take a `TimeString` and convert it to a `Date` object.
   * 
   * For now, just trust the server to send us valid dates.
  */
  toDate(t: TimeString): Date {
    return new Date(t)
  },
  /**
   * Take a `Date` object and convert it to a `TimeString`.
   */
  fromDate(d: Date): TimeString {
    return d.toISOString() as TimeString
  },
  /**
   * Take a `string` and convert it to a `TimeString`.
   * 
   * Returns `null` if date is invalid.
   */
  fromString(s: string): TimeString | null {
    const date = Date.parse(s)
    if (Number.isNaN(date)) {
      return null
    } else {
      return this.fromDate(new Date(date))
    }
  }
}
