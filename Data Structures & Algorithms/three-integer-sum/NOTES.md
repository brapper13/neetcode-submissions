# 3Sum — review

First encounter took 73 minutes, with hints. That is far over the
25-minute bar, but the two-pointer core went in fast. Almost the whole
overrun was the dedup layer, plus one crash. The problem goes on the
re-solve list for a cold timed run in about two weeks.

## How the submissions went

**The hashmap attempt (never submitted).** Run Two Sum with a map once
per anchor. The sums were correct, but duplicate triplets appeared and
could not be guarded away. Abandoned for the sort.

**Submission 3 — accepted, and crashes.** Sorted two-pointer with
`j = i + 1`, an anchor skip, and skip loops for both `j` and `k` after
a hit. The k-skip's guard was `k < len(nums)-1`. That is a right-edge
check on a loop that moves left, and nothing bounded it below.
`threeSum([]int{0, 0, 0, 0, 5})` panics with index out of range [-1].

NeetCode accepted it anyway. This is the third accepted-but-wrong
event, and the first that crashes instead of returning a bad answer.
It was caught by a local fuzz harness: 2,000 random arrays checked
against a brute-force reference.

**Submission 4.** The k-skip guard became `k > 0`. The panic is gone.

**Submission 5 — canonical.** The k-skip was deleted entirely, and the
j-skip guard tightened to `j < k`. Fuzz-verified: matches brute force
on 2,000 random inputs, with no duplicates and no panics.

## The dedup thought process

This is where the hour went. The wrong turns are the useful part.

**The duplicates appear.** The hashmap approach found correct sums
straight away. The first instinct was index guards: skip when `i == j`
or `i == k`. That instinct fails for a nameable reason. The guards
compare indices, but a duplicate triplet is two different index routes
to the same *values*. `{-1, 0, 1}` found via indices (0,1,2) and via
(4,1,2) passes every index test. No index check can see that they are
the same triplet.

**How bad is it?** There are two separate sources. Across anchors:
every triplet can be discovered from each of its three members, so
each one arrives three times. Within one anchor: if the array repeats
a value, the same pair is found again. `[-1, 1, 0, 0]` emits
`{-1, 0, 1}` twice from a single anchor pass.

**The filtering route.** You can dedup the output after the fact.
Sort each triplet's three values into a canonical form, then keep a
"seen" set. In Go the set key cannot be a slice, but `[3]int` is
comparable and works — the [[map-keys]] argument from go-notes. This
route is viable. It is also machinery: a set, a canonical form, and a
hash per candidate, all spent cleaning up duplicates the search
happily keeps generating.

**The reframe that unlocked it.** Instead of filtering duplicates out
of the output, make them impossible to construct. Impose a convention:
a triplet may only be discovered when the anchor is its *leftmost*
element in sorted order. Sort the array and start the pair search at
`j = i + 1`. Now each triplet has exactly one discovery path. The
cross-anchor duplicates are not filtered — they never exist. The
awkward index guards fall away unneeded.

**Sorting pays twice.** Equal values are now adjacent. The remaining
within-anchor duplicates become a local problem: skip past equal
anchor values, and after a hit, skip past equal `j` values. Each skip
is a one-line loop over neighbours instead of a set lookup.

The hour went on the first three steps — patching value-level
duplicates with index-level guards, then reaching for filtering
machinery, before questioning why the search was allowed to generate
duplicates at all. The transferable move became takeaway 18: constrain
generation before filtering results.

## Why the j-skip alone is enough

For a fixed anchor the target is fixed. So a pair is fully determined
by its `j` value — `nums[k]` must be `target - nums[j]`. The j-skip
guarantees the next hit has a strictly different `nums[j]`, which
forces a different pair. Different pair, same anchor: different
triplet. The k-skip was belt and braces. Interviewers ask "do you need
both skips?", and the answer is no.

## Skip-loop bounds

The live window is `(j, k)`. Values at or outside the window can never
be visited again, so every skip loop should stop at the window's edge:

```go
for j < k && nums[j] == nums[j+1] { j++ }
```

A guard phrased against `len(nums)` reads as a bounds patch, and
submission 3 showed how easily it goes wrong. Guards should state the
invariant, not the array edge.

## Other notes

- The anchor loop moved from `range` to a C-style `for`, because the
  anchor skip mutates `i`. A `range` loop would reset it each
  iteration.
- `slices.Sort` fails on NeetCode but the `min` builtin works, so the
  toolchain is at least Go 1.21. The judge pre-imports a fixed package
  set that lacks `slices`. Use `sort.Ints(nums)` — available
  everywhere.
