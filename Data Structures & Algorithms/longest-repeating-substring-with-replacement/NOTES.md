# Longest Repeating Character Replacement — review

About an hour to acceptance (submission 4), then fuzz-verified: 8,000
random string/k pairs against a brute-force reference, all matching.
Most of the hour went on replacing a wrong model, not on mechanics.

## The wrong model, and the probes that killed it

The first read was "the window may contain at most k *different*
characters", with a replacement budget that got spent on new
characters and refunded on... something. Two probes ended it:

`"AAAB"` and `"ABAB"`, both with k=1, both containing exactly two
distinct characters. The first is fixable with one replacement, the
second is not. Distinct-character counting cannot tell them apart, so
the invariant must involve *counts*.

And the choice of target letter answers itself: to make a window
uniform cheaply, keep its most frequent character and replace the
rest. So the cost of a window is `length - maxFreq`, and the window
is valid while that cost is at most k. No budget to spend or refund —
the cost is a property of the window's current contents.

## The placement bugs

The first coded attempt had the right pieces in the wrong places —
the same disease as Longest Substring, one problem later. The
entering character was only counted when it was already winning; the
shrink decremented the *arriving* character instead of the departing
`s[left]`; the invariant check hid inside a branch; and maxFreq was
decremented by guesswork. Trace `"ABB"` with k=0: the B's never get
counted, `count[B]` reaches -2, and the answer comes out 1 instead
of 2.

The fix was structural, not local: a straight-line loop body.

1. Count the arrival — unconditional.
2. `maxFreq = max(maxFreq, count[arrival])` — the only count that
   rose is the arrival's, so it is the only candidate.
3. One `if` for repair: on violation, decrement `count[s[left]]` and
   advance left.
4. Score the window — unconditional.

Nothing else is conditional, so there is no branch to owe bookkeeping
to.

## Jump versus creep

Longest Substring's left edge *jumps* — the map names the exact index
where the violation ends, and the skipped elements need no undo. Here
the left edge *creeps* one step per violation, because validity is an
aggregate (`length - maxFreq`) with no computable target index, and
every departing character must decrement its count on the way out.
The deciding question: what does the window have to undo when
something leaves? Nothing → you may jump. Per-element state → creep.

## The maxFreq subtlety — resolved

The shrink step lowers a count but never lowers `maxFreq`, so
`maxFreq` can overstate the current window's true max frequency and
the validity check turns lenient. The answer is still correct.

The reframe (found independently): `maxFreq` is a misnomer. The
counts are exact for the current window, but `maxFreq` is really "the
best frequency any window has contained so far" — a historical
maximum, just like the answer variable itself. The check was never
"this window against its own max frequency"; it is "this window's
size against the best certified size so far". Two pillars make that
legitimate:

**It cannot miss the best window.** A sub-window of a valid window is
valid: removing a character drops length by 1 and top frequency by at
most 1, so cost never rises. While the right edge sweeps the optimal
window with the left edge at or before its start, the true cost is
therefore ≤ k, and the lenient check passes even more easily. The
repair can never fire there, so the left edge cannot creep past the
optimal window's start, and the window reaches the optimal size.

**It cannot invent a size no valid window has.** A new larger size S
is only recorded when `S - maxFreq <= k`, so `S <= M + k`, where M
was — at the moment it was set — a genuine count: M copies of one
character actually present together. M identical characters plus k
replacements is the recipe for a real valid window of size M + k.
Every size the stale check waves through is achieved by some valid
window.

**The same fact removes the scoring `if`.** Each iteration the window
size grows by one (no repair) or stays flat (repair — the creep is
exactly one step). The size sequence is monotone non-decreasing, so
the final size is the historical maximum: delete `maxSeq` and its
`if`, and `return len(s) - left`. This depends on the repair being an
`if`, not a `while` — the one-step creep is what makes the sizes
monotone.

## Submissions 5 and 6 — the exercises, done

**Submission 5** applies the removal: `maxSeq` and its `if` deleted,
`return len(s) - left`, and the stale wrong comment removed with it.
Fuzz-verified.

**Submission 6** is the fixed-target alternative: for each of the 26
letters, one pass finding the longest window with at most k non-`c`
characters. The window state is a single `bad` counter — fixing the
target letter deletes the maxFreq machinery entirely, which makes this
version trivial to prove correct on the spot. The non-shrinking trick
composes into each pass: `len(s) - left` per letter, best of 26.
O(26n). Fuzz-verified.

Ranking for interviews: submission 5 is the fastest and hardest to
justify; submission 6 is barely slower and self-evident. Offer 6 if an
interviewer looks skeptical of the historical-maxFreq argument.

Toolchain note: `for i := range 26` compiled on NeetCode — their Go is
at least 1.22.
