# Trapping Rain Water — review

First encounter: 55 minutes, with hints. It's the block's Hard, so
that's a solid number. Accepted first submission, then fuzz-verified:
5,000 random arrays against a per-position brute force, all matching.

## How the solve went

- **First intuition (abandoned):** find wall pairs and compute each
  basin as `width * wallHeight - blocks`. Two problems. The skip loops
  had no window bounds — the exact 3Sum submission-3 bug, one day
  later. And the basin formula breaks on `[3, 0, 5, 0, 3]`: the 5
  pokes above the water line, and water doesn't wrap around a tower.
  A basin decomposition has to split at every bar taller than the
  water line, which fights the whole approach.
- **The reframe (hinted):** ask how much water sits on one single
  position instead of in one basin. Water at `k` is
  `min(maxLeft(k), maxRight(k)) - height[k]`.
- **The definition detail that matters:** make both maxes *inclusive*
  of position `k`. Then each side's max is at least `height[k]`, the
  formula can never go negative, and no clamp is needed. Exclusive
  maxes force a `max(0, ...)` wrapper.
- **From formula to O(n) (unassisted):** maxLeft and maxRight for every
  position are "everything left of k" and "everything right of k" —
  takeaway 5, the Product of Array Except Self move, with `max`
  instead of multiply. Two running passes, one combining pass. The
  transfer was recognised and executed without help.

## Polish (not resubmitted)

- The `if/continue/else` ladder in both passes is a hand-rolled `max`:
  `maxLeft[idx] = max(height[idx], maxLeft[idx-1])` — the 1.21
  builtin works on NeetCode.
- `trap([]int{})` panics on `maxLeft[0]`. LeetCode guarantees a
  non-empty array so the judge never sees it, but say the assumption
  out loud in an interview and offer the one-line guard.
- Two blank lines before the closing brace — gofmt would object.

## The O(1)-space follow-up

The prefix/suffix arrays exist only to answer `min(maxLeft, maxRight)`.
The two-pointer version keeps one running max per side and resolves
whichever side currently has the smaller max — that side's min is
already certain. Owed as a redo once the O(n) version has settled.
