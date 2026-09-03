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
