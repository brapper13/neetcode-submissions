# Review notes — longest-consecutive-sequence

**Verdict:** canonical O(n) shape reached. Set build, start gate on
!set[item-1], offset walk, running max. One real hole found (below).

## The arc

- Naive idea (chat): nested loops, O(n²).
- Second idea (chat): scan the value range min..max with a streak counter.
  Logically correct, but O(range) not O(n). His own pre-flight said
  "numbers large"; values reach ±10^9. [0, 10^9] runs a billion iterations
  for two elements. The bounded-domain pattern over-applied to an
  unbounded domain: its first loss after four wins.

## Why the range scan is bad, properly (pseudo-polynomial time)

Complexity is measured against input SIZE, and a number's size is its
written length. 10^9 occupies ~30 bits. A loop doing 10^9 steps for a
30-bit input is ~2^30 steps for 30 bits of input: exponential in input
length. It looks polynomial only because a value grows exponentially
faster than its digit count. Name: pseudo-polynomial.

Famous relative: knapsack's textbook DP is pseudo-polynomial in the
capacity, which is how knapsack stays NP-hard despite "having a DP".
Strong interview material; few candidates can explain it.

Precondition for the bounded-domain pattern, now explicit: before
iterating a value domain, compare its width to n. 26 letters, 9 digits,
freq <= n: scan. ±10^9: iterate elements, use a set.
- Final: gate-and-walk. The offset walk (set[item+seq]) is a clean
  formulation.

## The hole: iteration source breaks the complexity proof

The loop ranges over nums, which contains duplicates. Input: 50k copies of
1 + run 2..50001 → every copy passes the gate and walks 50k steps → O(n²).
The proof ("each run has one bottom") requires iterating UNIQUE values.
Fix: `for item := range set`. Lesson: where you loop from is part of the
design, not a detail.

## Do-over cleanup

- Stale comment describing the abandoned map-tracking design.
- Both early returns (len 0 and 1) are dead — natural flow handles them.

## Sudoku closure (same session)

Sub-16 added an own-words derivation of box = (i/3)*3 + j/3 as a comment.
Comprehension check passed.
