import type { Branded } from "@/ts_helpers/branding"
/**
 * `DurationNs` is the same as a number at runtime.
 * @see {@link Branded}
 */
export type DurationNs = Branded<number, 'duration_ns'>

const TIME_FACTORS = {
  NS_PER_MS:   1000,
  NS_PER_SEC:  1000 * 1000,
  NS_PER_MIN:  1000 * 1000 * 60,
  NS_PER_HOUR: 1000 * 1000 * 60 * 60,
  NS_PER_DAY:  1000 * 1000 * 60 * 60 * 24
}

export const DurationNs = {
  fromNanos(ns: number): DurationNs {
    return ns as DurationNs
  },
  fromMillis(ms: number): DurationNs {
    return ms * TIME_FACTORS.NS_PER_MS as DurationNs
  },
  fromSec(s: number): DurationNs {
    return s * TIME_FACTORS.NS_PER_SEC as DurationNs
  },
  fromMin(m: number): DurationNs {
    return m * TIME_FACTORS.NS_PER_MIN as DurationNs
  },
  fromHour(hr: number): DurationNs {
    return hr * TIME_FACTORS.NS_PER_HOUR as DurationNs
  },
  fromDay(d: number): DurationNs {
    return d * TIME_FACTORS.NS_PER_DAY as DurationNs
  },
  toNanos(ns: number): number {
    return ns
  },
  toMillis(ns: number): number {
    return ns / TIME_FACTORS.NS_PER_MS
  },
  toSec(ns: number): number {
    return ns / TIME_FACTORS.NS_PER_SEC
  },
  toMin(ns: number): number {
    return ns / TIME_FACTORS.NS_PER_MIN
  },
  toHour(ns: number): number {
    return ns / TIME_FACTORS.NS_PER_HOUR
  },
  toDay(ns: number): number {
    return ns / TIME_FACTORS.NS_PER_DAY
  }
}
