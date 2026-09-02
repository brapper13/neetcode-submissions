# When two pointers works, and when it cannot

The motivating puzzle: Two Sum on an unsorted array cannot be solved
with two pointers. Trapping Rain Water's input is also unsorted, and
two pointers solves it. So "the array must be sorted" is not the real
rule. This note is the real rule.

## The rule

Converging two pointers is a *pruning* argument over the space of
pairs. There are ~n²/2 candidate pairs `(i, j)`. The algorithm visits
at most n of them. So every single step must throw away a whole block
of pairs — everything the moved pointer is walking away from.

That throw-away is only legal with a proof. The proof is always the
same shape:

> Looking only at the current `(i, j)`, I can show that every
> remaining pair involving one of these two endpoints is a loser.
> So that endpoint is finished, and the pointer moves.

Two pointers works exactly when such a proof exists. The proof always
rests on something *monotone* — a quantity that only moves one way as
a pointer moves. Sortedness is one source of monotonicity. It is not
the only one.

## The same rule across the block

**Two Sum II (sorted).** The sum is too small. Moving `j` left only
shrinks it further, so `i` paired with *anything* remaining is also
too small. Every pair involving `i` is dead. `i++`. The sort is what
makes "too small now" imply "too small forever". This is the textbook
case: sortedness supplies the monotonicity directly.

**Two Sum (unsorted) — the failure case.** The sum is too small.
Which pointer moves? Nothing connects position to value, so a bigger
value could sit anywhere. No block of pairs can be proven dead, so
nothing can be discarded, and two pointers has no move to make. The
fix is to buy the missing information with memory instead: the hashmap
remembers every value seen, which is what order would have told you
for free. Memory and sortedness are two currencies for the same
purchase.

**Container With Most Water (unsorted — and it still works).** The
monotone quantity is the *width*: every move shrinks it. Area is
`width * min(wall, wall)`. Fix the shorter wall and consider its pairs
with nearer opposite walls: all narrower, and still capped by that
same shorter wall. Every one is provably no better than the pair just
scored. The shorter wall is finished. No sort anywhere — the proof
comes from the shrinking width plus the min-cap.

**Trapping Rain Water, O(1)-space version (unsorted).** The monotone
quantities are *manufactured*: a running max per side, and running
maxes only grow. Water at a position is capped by the smaller of the
two surrounding maxes. The side whose running max is currently smaller
already knows its cap for certain — the other side is guaranteed to
hold a wall at least that tall. So that side's position can be
resolved and stepped past. The proof rests on state the algorithm
built itself.

**Valid Palindrome — a different animal.** No pruning happens at all.
Every mirror pair must be checked, and the pointers simply enumerate
them. This is two pointers as *coverage*, not as search. Recognise the
difference: converging pointers that skip work need a proof;
converging pointers that just walk a symmetric structure do not.

**3Sum.** The input arrives unsorted and the pairing trick needs
order, so sort it. That is allowed because the answer wants values,
not indices — sorting destroys index information and nothing else.
When a problem wants indices (Two Sum classic), sorting costs you the
answer, and the hashmap is the way out.

## The checklist

Facing a new problem, ask in order:

1. What is the candidate space — pairs of what?
2. Standing at one candidate `(i, j)`, is there a rule that condemns
   every remaining pair touching one endpoint?
3. What monotone quantity makes that rule true? Sort order, a
   shrinking width, a growing running max — something must only move
   one way.
4. If no such quantity exists: can I create one? Sorting is legal when
   the answer survives it (values, not indices). Running state is
   legal always.
5. Still nothing? Then two pointers is the wrong tool. Pay with
   memory — a map of what has been seen — and expect O(n) space.

The one-line version: **two pointers is a proof that you may skip
work. No proof, no pattern.**

Related: takeaway 18 (constrain generation instead of filtering
output) is the same instinct applied to duplicates — both replace
cleanup with an argument about what can exist at all.
