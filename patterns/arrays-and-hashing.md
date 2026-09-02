# What a map actually buys you

The motivating puzzle this time: Two Sum has an obvious O(n²) answer —
check every pair. The map version is O(n) and visits each element
once. Where did a whole factor of n go? Nothing about the input
changed. What changed is that the loop stopped *searching* the past
and started *recalling* it.

## The rule

A map converts "scan everything I've seen for X" (O(n) per query)
into "look X up" (O(1) per query). That's the entire product. But it
comes with an obligation that is easy to miss:

> You must know the question before you store the answer. The key you
> insert *is* the question you'll ask later. Choose the wrong key and
> the map answers a question nobody asks.

So the design work in every hashing problem is the same: find the
question the loop will ask about its own past, phrase it as a key, and
store the answer as you go. The loop then never looks backwards.

## The rule across the block

**Contains Duplicate.** The question is "have I seen this exact value?"
The key is the value itself. The value stored doesn't matter, only
membership — which is why the Go set idiom is `map[T]struct{}`
(takeaway 8).

**Two Sum.** The question is not "have I seen `nums[i]`?" — it's
"have I seen `target - nums[i]`?" The key is the value, the stored
answer is the index, because the problem wants indices back. Notice
the question is phrased about the *complement*: you decide at insert
time what future lookups need to find.

**Valid Anagram.** The question is "how many of each letter?" The key
domain is 26 lowercase letters — small, dense, and known in advance.
When the key domain is bounded integers, an array outdoes the map:
`[26]int`, index arithmetic, no hashing at all (takeaway 1). An array
is a map whose keys are so well-behaved you can pre-allocate every
slot.

**Group Anagrams.** The question is "which group does this word belong
to?" The key must be something *all members share* — a canonical form.
Sort the word, or better, its `[26]int` count signature, which is
comparable and therefore a legal Go map key ([[map-keys]]). This is
the deepest key-design lesson in the block: sometimes the key is not a
value from the input but an equivalence class you compute. Choosing
the canonical form *is* the algorithm.

**Top K Frequent.** Two questions chained. First "how many times does
each value appear?" — a counting map. Then "which values have
frequency f?" — and since a frequency can't exceed n, the key domain
is bounded again, so the second structure is an array of buckets
indexed by frequency. Map for the unbounded domain, array for the
bounded one, each chosen by the same rule.

**Valid Sudoku.** The question is "has digit d appeared in this row /
column / box before?" That's 27 set-membership questions, all with
bounded domains: `[9][9]bool` three times. Visit order is irrelevant
because the tables just accumulate facts — the same order-independence
that makes all hashing solutions work.

**Longest Consecutive Sequence.** The question is "is `value - 1`
present?" — asked to detect the *start* of a run. A set answers it in
O(1), which is what lets an ordering problem dodge the O(n log n)
sort. The insight is not the set itself but noticing that "start of a
sequence" is a membership question in disguise.

**Product of Array Except Self.** The odd one out — no map at all. The
question "what's the product of everything left of me?" is answered by
*position*, not by value, so the structure is a running pass, not a
hash (takeaway 5). Worth keeping in the block's summary precisely
because it shows the boundary: maps answer questions keyed by value.
Questions keyed by position want prefix/suffix accumulation.

**Encode/Decode Strings.** Also not a hashing problem — it's
serialization, and its lesson (length-prefix framing, never in-band
markers) lives in its own notes. Included here only so the block list
is honest.

## What a map cannot buy

A map has no order. It cannot answer "what's the closest value below
X", "what came earlier", or "what's adjacent" — iteration order is
deliberately random in Go. When the question involves order, the
currencies are different: sort the input, or keep running state as
you walk it. That trade — memory versus order — is the bridge to the
two-pointer block: [two-pointers.md](two-pointers.md) closes with the
same coin from the other side.

## The checklist

1. Write down the question the loop wants to ask about elements it
   has already processed. Phrase it precisely — "have I seen the
   complement", not "have I seen stuff".
2. That question's subject is your key. What it needs back is your
   value. Membership only → `map[T]struct{}`.
3. Is the key domain bounded and dense (letters, digits, frequencies
   up to n)? Replace the map with an array.
4. Is the key shared by a whole group rather than one element? Compute
   a canonical form, and make sure it's a comparable type.
5. Does the question involve order or position? Stop — a map is the
   wrong tool. Running passes or sorting, per the two-pointer primer.

The one-line version: **a map turns search into recall, and the key
is the question — decide what you'll ask before you store the
answer.**
