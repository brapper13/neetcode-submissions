# 3Sum — review

First encounter: 73 minutes with hints. Way over the 25-minute bar, but
the two-pointer core went in fast — the whole overrun was the dedup
layer and one crash. On the re-solve list for a cold timed run in ~2
weeks.

## Progression

- **Pre-sub attempt (hashmap):** ran Two Sum with a map once per anchor.
  Correct sums, but duplicate triplets from two directions: the same
  triplet found from each of its three anchors, and again within one
  anchor pass when the array repeats values. Index guards can't fix a
  value-level problem. Abandoned for the sort.
- **submission-3 (accepted, crashes):** sorted two-pointer with
  `j = i + 1`, anchor skip, and both j- and k-skips after a hit. The
  k-skip guard was `k < len(nums)-1` — a right-edge check on a
  leftward-moving loop — and nothing bounded it below.
  `threeSum([]int{0, 0, 0, 0, 5})` panics with index out of range [-1].
  **NeetCode accepted it.** Third accepted-but-wrong event, and the
  first that crashes rather than returning a bad answer. Caught by a
  local fuzz harness: 2,000 random arrays diffed against a brute-force
  reference.
- **submission-4:** k-skip guard changed to `k > 0`. Panic gone.
- **submission-5 (canonical):** k-skip deleted entirely, j-skip guard
  tightened to `j < k`. Fuzz-verified: matches brute force on 2,000
  random inputs, no duplicates, no panics.

## The dedup thought process — where the hour went

Recorded blow by blow, because the wrong turns are the useful part.

**Step 1 — the duplicates appear.** The hashmap approach found correct
sums immediately. The first instinct was index guards: `i == j`,
`i == k`, skip. That instinct is wrong in a way worth naming: the
guards compare *indices*, but a duplicate triplet is two different
index routes to the same *values*. `{-1, 0, 1}` via indices (0,1,2)
and via (4,1,2) are both index-legal. No index test can see they're
the same triplet.

**Step 2 — how bad is it?** Two separate sources. Cross-anchor: every
triplet is discoverable from each of its three members, so each one
arrives three times. Within-anchor: if the array repeats a value, one
anchor pass finds the same pair again ([-1, 1, 0, 0] emits {-1, 0, 1}
twice from the same anchor).

**Step 3 — the filtering route.** You can dedup after the fact:
canonicalise each triplet by sorting its three values, then keep a
"seen" set. In Go the set key can't be a slice, but `[3]int` is
comparable and works ([[map-keys]] argument in go-notes). This route
is viable. It's also machinery: a set, a canonical form, hashing per
candidate — all spent *cleaning up* duplicates the search happily
keeps generating.

**Step 4 — the reframe that unlocked it.** Instead of filtering
duplicates out of the output, make them impossible to construct.
Impose a convention: a triplet may only be discovered when the anchor
is its *leftmost* element in sorted order. Sort the array and start
the pair search at `j = i + 1`. Now each triplet has exactly one
discovery path — cross-anchor duplicates don't get filtered, they
never exist. The awkward index guards from step 1 fall out unneeded.

**Step 5 — sorting pays twice.** Equal values are now adjacent, so
the remaining (within-anchor) duplicates become a *local* problem:
skip past equal anchor values, and after a hit, skip past equal `j`
values. Each skip is a one-line loop over neighbours instead of a
global set lookup.

The hour went on steps 1–3: trying to patch value-level duplicates
with index-level guards, then reaching for filtering machinery, before
questioning why the search was allowed to generate duplicates at all.
The generalisable move: when output needs deduplicating, first ask
whether the search space can be shaped so duplicates are
unrepresentable. Constrain generation before filtering results.

## Why the j-skip alone dedups

For a fixed anchor the target is fixed, so a pair is fully determined
by its `j` value — `nums[k]` must be `target - nums[j]`. The j-skip
guarantees the next hit has a strictly different `nums[j]`, which
forces a different pair. The k-skip was belt and braces; interviewers
ask "do you need both?", and the answer is no.

## Skip-loop bounds

The live window is `(j, k)`. Values at or outside the window can never
be visited again, so every skip loop should stop at the window's edge:

```go
for j < k && nums[j] == nums[j+1] { j++ }
```

A guard phrased against `len(nums)` reads as a bounds patch and, as
submission-3 showed, is easy to get wrong. Guards should state the
invariant, not the array edge.

## Other notes

- Switched the anchor loop from `range` to a C-style `for` because the
  anchor skip mutates `i` — a `range` loop would reset it each
  iteration.
- `slices.Sort` fails on NeetCode but the `min` builtin works
  (max-water-container proved it), so the toolchain is ≥1.21 — the
  judge just pre-imports a fixed package set without `slices`. Use
  `sort.Ints(nums)`; it's available everywhere.
