# Problem playbook — Arrays & Hashing

Built from the nine problems of the first block. Covers only techniques
already used. Re-read before timed sessions.

## The ritual

**Before writing code**
1. Write the constraint comments: input sizes, value ranges, empty cases.
2. Ask: what input breaks the obvious idea? (duplicates, zeros, ties,
   empty, single element, the marker inside the data)
3. Ask: does the problem constrain the METHOD? ("without division",
   "in one pass", "O(1) space")
4. Say the brute force and its complexity, then look for the O(n) shape.

**Choosing the technique**

**"Have I seen this before?" / "does X exist?"**
Use a set. A `map[T]bool` set can be read with plain `if set[x]`, because a
stored `false` never exists — the zero value always means absent. If the map
stores real data (an index, a count), a stored 0 is legitimate, so read with
`v, ok := m[k]` instead. ([Takeaway 7](TAKEAWAYS.md#7-comma-ok-when-zero-could-be-real-data-direct-indexing-when-it-cannot).)

**"Count occurrences of each thing."**
A map counter by default. When the things come from a small bounded range
(26 letters, 9 digits), use a fixed array like `[26]int` instead — no
hashing, no allocation. ([Takeaway 1](TAKEAWAYS.md#1-when-the-keys-are-small-bounded-integers-use-an-array-instead-of-a-map).)

**"Group items that share an identity."**
Build a canonical form for each item and use it as a map key. If the
canonical form is a bounded-int count, the count array itself can be the
key, because arrays are comparable. ([Takeaway 9](TAKEAWAYS.md#9-what-can-and-cannot-be-a-go-map-key).)

**"Each answer combines everything left of i with everything right of i."**
Two running passes: forward for the left part, backward for the right part.
A single running variable per pass gives O(1) extra space.
([Takeaway 5](TAKEAWAYS.md#5-everything-left-of-i-plus-everything-right-of-i-means-two-running-passes).)

**"Top k by frequency."**
Count first, then bucket numbers by their frequency. A frequency can never
exceed `len(nums)`, so a slice indexed by frequency replaces sorting.
([Takeaway 4](TAKEAWAYS.md#4-derive-the-bound-out-loud-then-pick-the-container).)

**"Find runs / sequences in unordered data."**
Put everything in a set. Only start walking from a value whose predecessor
is absent (`!set[x-1]`), then walk upward while successors exist. Iterate
the set, not the input, or duplicates break the O(n) argument.
([Takeaway 6](TAKEAWAYS.md#6-count-what-your-loop-actually-iterates).)

**"Pack arbitrary data into one string and get it back exactly."**
Write each piece's length in front of it, and read lengths at known
positions when decoding. Never search the data for a marker — the data can
contain the marker.

**Two complexity checks before submitting**
1. What does the loop iterate? Elements is O(n). Values is O(range) —
   pseudo-polynomial. Only scan a value domain when its width is
   comparable to n (26 letters, 9 digits, freq ≤ n). ±10^9 is not.
2. What allocates inside the hot loop? Maps and Sprintf in a loop were
   the real cost every time it mattered. Function calls were never.

**Design smells learned the hard way**
- Guards piling up → look for the shape that makes the bad state
  unrepresentable (check-before-insert killed the self-pair guard AND the
  ordering swap).
- Clever control flow that needs a trace → replace with nested loops that
  don't (the box odometer).
- Repeated index expressions → extract a named local (`d :=`), it's where
  swap bugs live.

**Before every paste**
- grep Println
- delete dead code, stale comments, debug leftovers
- gofmt (write locally, paste after)

## Conversions (used constantly in this block)

| From → to | How | Why it works |
|---|---|---|
| digit char → int | `int(x - '0')` | digit code points are consecutive |
| letter → 0..25 index | `int(x - 'a')` | a..z code points are consecutive |
| numeric string → int | `strconv.Atoi(s)` | |
| int → string | `strconv.Itoa(n)` | not Sprintf |
| 0..25 index → letter | `byte('a' + i)` | same trick, reversed |
| int digit → char | `byte('0' + d)` | |

`s[i]` gives a byte. `range s` gives runes. The subtraction trick works on
either, but only under an ASCII constraint — state the constraint.

## Do-over session (first run: 2026-08-23)

Timed, from blank, notes closed. Record minutes per problem.
1. is-anagram — target: canonical single map or [26]int directly
2. two-sum — target: one-pass check-before-insert directly
3. top-k — target: buckets directly, pre-sized map
4. encode-decode — target: clean length-prefix, Builder in encode,
   slices in decode, no Println
5. valid-sudoku — single pass, d extracted, combined checks
6. longest-consecutive — set iteration, dead returns removed

Gate reference (week 4): easies reliably ≤ 15 minutes.
