# Permutation in String — review

Accepted at submission 7, rewritten properly as submission 8. Both
fuzz-verified against brute force on 10,000 random pairs.

## Submission 7 — correct, but not a window

For every start position, copy the target counts and re-scan the full
window from scratch. O(n·m), accepted only because the constraints
are small. Written mid-block, one day after three sliding windows —
a pattern miss worth recording: the technique didn't transfer on its
own, and an interviewer's "can you do better?" would have landed
immediately.

Two Go touches worth keeping from it: `local = countS1` copies a
[26]int by value (the assignment slices would betray — see
[[map-keys]]), and `continue Outer` with a label is legitimate,
rarely-seen Go.

## Submission 8 — lockstep

The unlock question: the windows at positions i and i+1 differ by
exactly two characters — one enters at the front, one falls off the
back. So one running count array, adjusted twice per step, and an
array compare (`countS1 == local`, O(26), legal because arrays are
comparable values) whenever the window is at full size.

This is the third window species, completing the set:

- **Jump** (Longest Substring): left teleports to a computed
  violation boundary; skipped elements need no undo.
- **Creep** (Character Replacement): left steps one at a time,
  unwinding aggregate state per departing element.
- **Lockstep** (here): left is not a decision at all — the window
  size is pinned by the problem, so left is arithmetic
  (`right - len(s1) + 1`), no invariant, no repair branch.

The deciding question for lockstep: does the problem fix the answer's
size in advance? A permutation of s1 has exactly len(s1) characters —
no other window size is interesting.

## Polish (not resubmitted)

- The `else` after `return true` is dead structure — a return needs
  no else.
- The "change to map" comment is stale — it describes a plan that
  became two arrays.

## Open

- Attempts 1–6 unrecovered: what they died of is unrecorded (asked;
  takeaway 14 wants it if the pastes still exist).
- Untimed.
