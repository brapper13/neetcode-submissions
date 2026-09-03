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

## The O(1)-space version (submission 1)

Done the next day, derived rather than looked up, from three guided
questions. Fuzz-verified against the same brute-force reference.

The derivation in full, because it is the defensible part.

The two arrays in submission 0 exist only to answer
`min(maxLeft(k), maxRight(k))`. With pointers at both ends carrying
running maxes, standing at position `i` you know `maxLeft(i)` exactly
(that side is fully scanned) but only a *lower bound* on
`maxRight(i)` — the unscanned middle could raise it. A lower bound on
one argument of `min` is enough exactly when the other argument is
smaller. So when `maxL <= maxR`, position `i`'s water is certain:
`maxL - height[i]`. Pay it and move `i`. Symmetric on the other side.

In the primer's language: the running maxes are the monotone
quantities (they only grow), and the proof of safe skipping is that a
growing lower bound can never fall below a value it already exceeds.

**The meeting cell needs no payment.** The loop pays each cell as a
pointer leaves it, so the one cell where the pointers meet is never
paid. That is safe because the meeting cell is always a global
maximum. A pointer only stalls on a cell that beats the other side's
running max. If anything taller sat in the walker's remaining
territory, the walker's max would overtake and the roles would flip —
so neither pointer can be parked on less than the tallest bar when
the other arrives. For a global maximum both inclusive side-maxes
equal its own height, and the water formula gives exactly 0.

Polish on submission 1:

- The two branches are guarded by complementary `if`s instead of
  `if/else`. It works only because neither branch touches the maxes;
  an edit inside the first branch could make both fire in one
  iteration. `else` states the exclusivity instead of re-deriving it
  (takeaway 3).
- Same empty-array panic as submission 0, same verdict: guaranteed
  away by the constraints, worth saying aloud.
