# Longest Substring Without Repeating Characters — review

This one has the longest arc of any problem so far. It was quiz
question P1 before it was a solve, and the quiz work and the solve
tell one story.

## The arc

**Quiz round one (cold):** right instinct — scan with a seen-map —
but no technique name, and the sketch reset everything on a duplicate.
Homework assigned: trace the reset procedure on `"dvdf"`.

**Quiz round two:** the trace done correctly (returns 2, truth is 3),
the diagnosis correct (the reset throws away the still-valid
characters between the two occurrences), technique named, and the map
upgraded from bool to last-seen index. That earned the classic trap of
the index-jump variant: `"abba"`, where a stale map entry tells the
left edge to move *backwards*. The guard was derived before any code
was written: the left edge only moves forward,
`left = max(left, prev+1)`.

**Submissions up to 6 (accepted):** correct, fuzz-verified — but it
took six attempts, and the shape explains why. Three branches (fresh,
stale duplicate, live duplicate), each carrying its own copy of the
count-and-max bookkeeping. Keeping three copies consistent is what
the failed attempts were spent on. Worse, staleness was defended
twice: the `max` guard already neutralises stale entries, yet an
explicit stale branch detected them separately. Two redundant
defenses, triple the paths — takeaway 3 in action.

**Submissions 9–10 (canonical):** produced by answering one question:
what is `currSeq` equal to at every point, in terms of `idx` and
`left`? It is always `idx - left + 1`. So the variable holds nothing
the pointers don't already know, and once it's deleted the branches
collapse: update `left` on a duplicate, record the index, take the
max. One path through the loop. Fuzz-verified against brute force on
8,000 random strings over a 4-letter alphabet (small alphabet =
maximum duplicate pressure) plus `"abba"`, `"dvdf"`, `"tmmzuxt"`.

## What each failed attempt died of (takeaway 14)

Recovered after the fact — the failures were pasted back for autopsy.

- **Attempt 1:** the jump branch paid none of its bookkeeping debts —
  no index recorded, count ignored `left`, no max update.
- **Attempt 2:** stale branch added; count off by one
  (`idx - left`, missing the `+1`); jump branch still not recording
  the index.
- **Attempt 3:** off-by-one fixed. The unrecorded index now acts
  alone: on `"aabaa"` the map points at a pre-jump occurrence, the
  next duplicate is misclassified as stale, and a window with two
  `a`s gets counted. The defensive stale branch became the vehicle
  for the wrong answer.
- **Attempt 4:** index recording fixed — and the count regressed,
  moved above the `left` update and minus its `+1` again. Fixing one
  piece of redundant state destabilised another.
- **Attempt 5:** byte-identical resubmit of attempt 4. A burned
  attempt. New checklist item: when a fix fails, diff against the
  previous paste before submitting again.
- **Attempts 7–8 (after the first acceptance, during the collapse):**
  `currSeq` deleted, branches merged — but the max update stayed
  *inside* the fresh branch. Any window whose best extent ends on a
  duplicate character is never measured. The fuzzer kills it with
  `"abbdcacabb"` (returns 3, truth 4); `"tmmzuxt"` is the classic
  form of the same shape. Fixed in submission 9 by making the max
  check unconditional, outside all branches: score the window every
  iteration, no exceptions.

Seven failures, zero algorithm errors. The `abba` guard was correct
before any code existed. Every death was bookkeeping placement —
which branch owes which update — and the final form wins by having
one branch and an unconditional score, so the question cannot arise.

## Remaining polish (not resubmitted)

- Both branches end in `seen[char] = idx`, so the assignment hoists
  out and the `else` disappears:

  ```go
  if prev, ok := seen[char]; ok {
  	left = max(left, prev+1)
  }
  seen[char] = idx
  ```

- `>=` in the max check can be `>` — equality changes nothing.
- `range s` yields byte offsets, so `idx - left + 1` measures bytes.
  Equal to character count under the ASCII constraint; silently wrong
  for multibyte input (takeaway 11). State the assumption aloud.

## The transferable lesson

The window's two edges each move only one direction. That monotone
property is simultaneously the O(n) argument (each index enters and
leaves the window once), the fix for `"abba"` (a forward-only left
edge cannot obey a stale entry), and the reason no bookkeeping
variable is needed (the window *is* the count). When a sliding-window
solution is fighting you, check whether some state you added is
duplicating what the pointers already encode.
