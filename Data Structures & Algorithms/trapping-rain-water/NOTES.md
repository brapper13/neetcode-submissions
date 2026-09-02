# Trapping Rain Water — review

First encounter took 55 minutes, with hints at two points. This is the
block's Hard, so that is a solid number. The first submission was
accepted. It was then fuzz-verified locally: 5,000 random arrays
checked against a brute-force reference, all matching.

## The first idea, and why it died

The starting intuition was basin-based. Find a wall on the left and a
wall on the right, then compute the water between them as
`width * wallHeight - blocks`.

Two things killed it.

The sketch had skip loops with no window bounds. `height[j-1]` had
nothing stopping `j` from reaching 0. This is the same bug as 3Sum
submission-3, one day later. Every skip loop gets a window bound.

The basin formula itself breaks on `[3, 0, 5, 0, 3]`. The 5 in the
middle pokes above the water line between the two 3s. Water does not
wrap around a tower — the tower splits the pool in two. A basin
approach would have to split at every bar taller than the water line,
and at that point the approach is fighting you.

## The reframe

Stop asking how much water a basin holds. Ask how much water sits on
top of one single position. The answer for position `k`:

```
water(k) = min(maxLeft(k), maxRight(k)) - height[k]
```

One definition detail carries real weight. Both maxes must *include*
position `k` itself. Then each side's max is at least `height[k]`, so
the formula can never go negative, and no clamp is needed. With
exclusive maxes, the tallest bar produces a negative and everything
needs a `max(0, ...)` wrapper.

## From formula to O(n)

The formula needs `maxLeft` and `maxRight` for every position. That is
"everything left of k" and "everything right of k" — takeaway 5, the
same move as Product of Array Except Self, with `max` in place of
multiply. One forward pass fills `maxLeft`, one backward pass fills
`maxRight`, one final pass sums the water.

This transfer was recognised and executed without help. Best moment of
the solve.

## Polish for next time

None of this was worth a resubmit.

- The `if/continue` ladder in each pass is a hand-rolled `max`. One
  line does it: `maxLeft[idx] = max(height[idx], maxLeft[idx-1])`.
- `trap([]int{})` panics on `maxLeft[0]`. LeetCode guarantees a
  non-empty array, so the judge never sees it. In an interview, say
  the assumption out loud and offer the one-line guard.
- Two blank lines sit before the closing brace. gofmt would remove
  them.

## Owed: the O(1)-space version

The two arrays exist only to answer `min(maxLeft, maxRight)` at each
position. The two-pointer version drops them. Keep one running max per
side, and resolve whichever side currently has the smaller max — that
side's `min` is already certain, because the other side guarantees a
wall at least as tall. Do this as a redo once the O(n) version has
settled.
