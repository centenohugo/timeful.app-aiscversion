import { describe, expect, it } from "vitest"
import {
  AUTO_SCROLL_EDGE_SIZE,
  AUTO_SCROLL_MAX_SPEED,
  exceedsColumnLockThreshold,
  formatSlotRangeLabel,
  getAutoScrollDelta,
} from "./touch_drag_utils"

const TOP = 100
const BOTTOM = 700

describe("getAutoScrollDelta", () => {
  it("does not scroll when the finger is away from the edges", () => {
    expect(getAutoScrollDelta(400, TOP, BOTTOM)).toBe(0)
    expect(getAutoScrollDelta(TOP + AUTO_SCROLL_EDGE_SIZE, TOP, BOTTOM)).toBe(0)
    expect(
      getAutoScrollDelta(BOTTOM - AUTO_SCROLL_EDGE_SIZE, TOP, BOTTOM)
    ).toBe(0)
  })

  it("scrolls up near the top and down near the bottom", () => {
    expect(getAutoScrollDelta(TOP + 10, TOP, BOTTOM)).toBeLessThan(0)
    expect(getAutoScrollDelta(BOTTOM - 10, TOP, BOTTOM)).toBeGreaterThan(0)
  })

  it("speeds up the deeper the finger goes into the edge zone", () => {
    const shallow = getAutoScrollDelta(BOTTOM - 40, TOP, BOTTOM)
    const deep = getAutoScrollDelta(BOTTOM - 5, TOP, BOTTOM)
    expect(deep).toBeGreaterThan(shallow)
  })

  it("caps the speed when the finger is past the edge", () => {
    expect(getAutoScrollDelta(TOP - 200, TOP, BOTTOM)).toBe(
      -AUTO_SCROLL_MAX_SPEED
    )
    expect(getAutoScrollDelta(BOTTOM + 200, TOP, BOTTOM)).toBe(
      AUTO_SCROLL_MAX_SPEED
    )
  })

  it("does not scroll when the visible area is too small", () => {
    expect(getAutoScrollDelta(110, 100, 200)).toBe(0)
  })
})

describe("exceedsColumnLockThreshold", () => {
  it("keeps the column locked for small sideways jitter", () => {
    expect(exceedsColumnLockThreshold(20, 100)).toBe(false)
    expect(exceedsColumnLockThreshold(-59, 100)).toBe(false)
  })

  it("unlocks once the finger travels most of a column in either direction", () => {
    expect(exceedsColumnLockThreshold(60, 100)).toBe(true)
    expect(exceedsColumnLockThreshold(-80, 100)).toBe(true)
  })

  it("never unlocks before the timeslot size is known", () => {
    expect(exceedsColumnLockThreshold(500, 0)).toBe(false)
  })
})

describe("formatSlotRangeLabel", () => {
  it("shows a single day once", () => {
    expect(
      formatSlotRangeLabel({
        startDay: "Mon",
        endDay: "Mon",
        startTime: "9 am",
        endTime: "11:30 am",
      })
    ).toBe("Mon · 9 am – 11:30 am")
  })

  it("shows a span of days", () => {
    expect(
      formatSlotRangeLabel({
        startDay: "Mon",
        endDay: "Wed",
        startTime: "9:00",
        endTime: "10:00",
      })
    ).toBe("Mon – Wed · 9:00 – 10:00")
  })
})
