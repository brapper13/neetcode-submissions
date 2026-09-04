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

## The maxFreq subtlety

The shrink step lowers a count but never lowers `maxFreq`, so
`maxFreq` can overstate the window's true max frequency and the
validity check turns lenient. The answer is still correct — the
argument is recorded here after being worked through, see the
conversation-notes section below.

## Open items

- The comment block still carries the abandoned "at most k different
  characters" model — a stale comment lying above correct code.
  Delete on next touch.
- The maxFreq argument (above) to be written up once answered.
