# Review notes — products-of-array-discluding-self (submissions 1–2)

**Verdict:** sub-2 is the correct two-pass prefix/suffix solution. O(n), no
zero special-casing. Sub-1 was the division solution with complete zero
handling — correct but forbidden by the problem statement.

## The arc

- Naive (shared in chat, not submitted): O(n²) pairwise. Would TLE at 100k.
- Sub-1: division + full zero case analysis (0/1/2+ zeros all handled).
  Interview response would be "good, now without division".
- Sub-2: left product × right product, one forward pass, one backward pass.
  The case tree vanished. Zeros need nothing.

## The catch worth remembering

preProducts and postProducts are maps keyed by 0..n-1. The keys ARE array
indices. TAKEAWAYS entry 1 (bounded dense int keys → indexed slice) did not
fire on its purest case. Swap both for make([]int, len(nums)). Lesson: the
pattern is learned but not yet automatic under new-problem pressure. The
do-over exists for exactly this.

## Follow-up: CLOSED (submission 3)

Reached the optimal form unprompted past the hint: both directions run on a
single variable, the output array receives left products then absorbs right
products in the backward pass. O(n) time, O(1) extra space. Also retired
both recurring style notes (dead underscore, redundant cap) and the maps in
the same submission. Problem complete at the highest level.

## Recurring style (third appearance each — retire these)

- `for i, _ := range` → `for i := range`.
- `make([]int, n, n)` → `make([]int, n)`.

## Credit

- Design explained in own words in a comment before the code.
- Backward pass indexing exactly right first time.
- Pre-flight includes the zero analysis that motivated the final design.
