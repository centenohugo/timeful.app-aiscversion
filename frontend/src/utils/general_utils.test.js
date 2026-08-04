import { describe, expect, it } from "vitest"
import {
  canScheduleEvent,
  isEventOwner,
  isEventScheduler,
  isGuestEvent,
} from "./general_utils"
import { guestUserId } from "../constants"

const OWNER = "507f1f77bcf86cd799439011"
const DELEGATE = "507f1f77bcf86cd799439012"
const BYSTANDER = "507f1f77bcf86cd799439013"

const ownedEvent = (schedulers) => ({
  ownerId: OWNER,
  ...(schedulers === undefined ? {} : { schedulers }),
})

describe("isGuestEvent", () => {
  it("is true only for events created without signing in", () => {
    expect(isGuestEvent({ ownerId: guestUserId })).toBe(true)
    expect(isGuestEvent(ownedEvent())).toBe(false)
  })

  it("tolerates a missing event", () => {
    expect(isGuestEvent(undefined)).toBe(false)
  })
})

describe("isEventOwner", () => {
  it("matches the owner and nobody else", () => {
    expect(isEventOwner(ownedEvent(), OWNER)).toBe(true)
    expect(isEventOwner(ownedEvent(), BYSTANDER)).toBe(false)
  })

  it("is false for a signed out visitor", () => {
    // A signed out visitor has no id, which must not match a missing ownerId
    expect(isEventOwner({}, undefined)).toBe(false)
  })
})

describe("isEventScheduler", () => {
  it("is false when the owner has delegated to nobody", () => {
    expect(isEventScheduler(ownedEvent(), DELEGATE)).toBe(false)
    expect(isEventScheduler(ownedEvent([]), DELEGATE)).toBe(false)
  })

  it("is true only for users in the schedulers list", () => {
    const event = ownedEvent([DELEGATE])
    expect(isEventScheduler(event, DELEGATE)).toBe(true)
    expect(isEventScheduler(event, BYSTANDER)).toBe(false)
  })

  it("handles the null the API sends when the field is unset", () => {
    expect(
      isEventScheduler({ ownerId: OWNER, schedulers: null }, DELEGATE)
    ).toBe(false)
  })

  it("is false for a signed out visitor", () => {
    expect(isEventScheduler(ownedEvent([DELEGATE]), undefined)).toBe(false)
  })
})

describe("canScheduleEvent", () => {
  it("lets the owner schedule their own event", () => {
    expect(canScheduleEvent(ownedEvent(), OWNER)).toBe(true)
  })

  it("blocks a respondent who has not been delegated to", () => {
    expect(canScheduleEvent(ownedEvent([DELEGATE]), BYSTANDER)).toBe(false)
  })

  it("lets a delegate schedule", () => {
    expect(canScheduleEvent(ownedEvent([DELEGATE]), DELEGATE)).toBe(true)
  })

  it("keeps guest events open to everybody, as they were before delegation", () => {
    const event = { ownerId: guestUserId }
    expect(canScheduleEvent(event, BYSTANDER)).toBe(true)
    expect(canScheduleEvent(event, undefined)).toBe(true)
  })

  it("blocks a signed out visitor on an owned event", () => {
    expect(canScheduleEvent(ownedEvent([DELEGATE]), undefined)).toBe(false)
  })
})
