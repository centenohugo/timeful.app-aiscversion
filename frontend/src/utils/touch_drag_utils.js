/*
  Touch drag utils — pure helpers for selecting timeslots with a finger
*/

/** Height (px) of the zone at the top/bottom of the visible grid that triggers auto-scroll */
export const AUTO_SCROLL_EDGE_SIZE = 56
/** Max auto-scroll speed in px per animation frame */
export const AUTO_SCROLL_MAX_SPEED = 14
/** Fraction of a column's width the finger must travel sideways before a drag spans multiple days */
export const COLUMN_LOCK_THRESHOLD = 0.6

/**
 * Returns how many px to scroll this frame given the finger's clientY and the visible
 * bounds of the grid. Negative scrolls up, positive scrolls down, 0 means don't scroll.
 * Speed ramps up linearly the deeper the finger goes into the edge zone.
 */
export const getAutoScrollDelta = (
  pointerY,
  topBound,
  bottomBound,
  edgeSize = AUTO_SCROLL_EDGE_SIZE,
  maxSpeed = AUTO_SCROLL_MAX_SPEED
) => {
  // Visible area too small to have distinct top and bottom zones
  if (bottomBound - topBound <= edgeSize * 2) return 0

  const speedForDepth = (distanceFromEdge) => {
    const depth = Math.min(edgeSize - distanceFromEdge, edgeSize)
    return Math.ceil((depth / edgeSize) * maxSpeed)
  }

  const distanceFromTop = pointerY - topBound
  if (distanceFromTop < edgeSize) return -speedForDepth(distanceFromTop)

  const distanceFromBottom = bottomBound - pointerY
  if (distanceFromBottom < edgeSize) return speedForDepth(distanceFromBottom)

  return 0
}

/** Whether a touch drag has moved far enough sideways to start selecting across days */
export const exceedsColumnLockThreshold = (
  dx,
  colWidth,
  threshold = COLUMN_LOCK_THRESHOLD
) => {
  return colWidth > 0 && Math.abs(dx) >= colWidth * threshold
}

/** Formats a selected range, e.g. "Mon – Wed · 9 am – 11:30 am" */
export const formatSlotRangeLabel = ({
  startDay,
  endDay,
  startTime,
  endTime,
}) => {
  const days = startDay === endDay ? startDay : `${startDay} – ${endDay}`
  return `${days} · ${startTime} – ${endTime}`
}
