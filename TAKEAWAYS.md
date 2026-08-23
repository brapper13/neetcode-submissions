# Rolling takeaways

Lessons from the problems done so far, one section each. Every entry states
the rule in full sentences, names the problems that taught it, and links to
the exact lines of code. Each problem name links to a STATEMENT.md in its
folder with the full problem description, an example, and the constraints
— so you can re-anchor without opening NeetCode.

Re-read this before the block do-over and before any timed session.

## Patterns

### 1. When the keys are small bounded integers, use an array instead of a map

A map costs hashing and allocation. If the keys can only be a small known
range (26 letters, frequencies up to `len(nums)`), an array indexed by the
key does the same job with none of that cost.

The precondition matters: the range must be small. If values can span
billions, an array over the range is worse than the map (see takeaway 6).

Seen in:
- **[Valid Anagram](Data%20Structures%20&%20Algorithms/is-anagram/STATEMENT.md)** — given two strings, do they contain the same letters
  the same number of times?
  → [is-anagram/submission-9.go, lines 5–8](Data%20Structures%20&%20Algorithms/is-anagram/submission-9.go#L5-L8):
  `[26]int` counter, one `++` and one `--` per position.
- **[Group Anagrams](Data%20Structures%20&%20Algorithms/anagram-groups/STATEMENT.md)** — group a list of strings so anagrams sit together.
  → [anagram-groups/submission-4.go, line 3](Data%20Structures%20&%20Algorithms/anagram-groups/submission-4.go#L3):
  the `[26]int` array itself is the map key, so no sorting is needed.
- **[Top K Frequent](Data%20Structures%20&%20Algorithms/top-k-elements-in-list/STATEMENT.md)** — return the k most common numbers in a list.
  → [top-k-elements-in-list/submission-5.go, line 9](Data%20Structures%20&%20Algorithms/top-k-elements-in-list/submission-5.go#L9):
  buckets indexed by frequency, sized `len(nums)+1` because a number
  cannot appear more often than the list is long.
- **[Valid Sudoku](Data%20Structures%20&%20Algorithms/valid-sudoku/STATEMENT.md)** — does a 9×9 board repeat a digit in any row, column,
  or 3×3 box?
  → [valid-sudoku/submission-16.go, line 3](Data%20Structures%20&%20Algorithms/valid-sudoku/submission-16.go#L3):
  `[9][9]bool` seen-tables instead of maps, because both indices are
  bounded by 9.

### 2. Order the loop so bad states cannot exist

In one-pass Two Sum, you look the complement up *before* inserting the
current number. The current number is never in the map when you check, so
pairing a number with itself is impossible. No guard needed.

The general move: when you catch yourself writing a guard like
`if i == j { continue }`, ask whether a different processing order makes
the bad state unrepresentable instead.

Seen in:
- **[Two Sum](Data%20Structures%20&%20Algorithms/two-integer-sum/STATEMENT.md)** — find the two indices whose values add up to a target.
  → [two-integer-sum/submission-4.go, lines 4–7](Data%20Structures%20&%20Algorithms/two-integer-sum/submission-4.go#L4-L7):
  the lookup on line 4 runs before the insert on line 7.
- Counterexample: the pairwise Group Anagrams attempt needed exactly that
  guard →
  [anagram-groups/submission-3.go, line 12](Data%20Structures%20&%20Algorithms/anagram-groups/submission-3.go#L12).

### 3. Simplifying control flow kills latent bugs

The bucket-walking loop in Top K first used two hand-managed counters
(`m` and `index`). Rewriting it as a plain nested range with an early
return was one edit, and that edit was also the bug fix. When a loop needs
bookkeeping variables you have to think hard about, the bug is usually
already in there.

Seen in:
- **[Top K Frequent](Data%20Structures%20&%20Algorithms/top-k-elements-in-list/STATEMENT.md)** — return the k most common numbers in a list.
  Before: [top-k-elements-in-list/submission-2.go, lines 23–34](Data%20Structures%20&%20Algorithms/top-k-elements-in-list/submission-2.go#L23-L34).
  After: [top-k-elements-in-list/submission-3.go, lines 23–30](Data%20Structures%20&%20Algorithms/top-k-elements-in-list/submission-3.go#L23-L30).

### 4. Derive the bound out loud, then pick the container

A number cannot appear more often than the list is long, so every
frequency is at most `len(nums)`. That derivation is what justifies a
bucket array of size `len(nums)+1` instead of sorting the frequencies.
In an interview, say the derivation — it is the part that scores.

Seen in:
- **[Top K Frequent](Data%20Structures%20&%20Algorithms/top-k-elements-in-list/STATEMENT.md)** —
  [top-k-elements-in-list/submission-5.go, line 6](Data%20Structures%20&%20Algorithms/top-k-elements-in-list/submission-5.go#L6)
  is the derivation written as a comment, and
  [line 9](Data%20Structures%20&%20Algorithms/top-k-elements-in-list/submission-5.go#L9)
  is the container sized from it.

### 5. "Everything left of i" plus "everything right of i" means two running passes

When `answer[i]` combines information about all elements before `i` and
all elements after `i`, build it with a forward pass and a backward pass.
Neighbouring positions differ by exactly one operation, so each pass needs
only a single running variable and the whole thing is O(n) time, O(1)
extra space.

Compare the division approach, which needed a case tree for zeroes. This
shape needed no special cases at all.

Seen in:
- **[Product of Array Except Self](Data%20Structures%20&%20Algorithms/products-of-array-discluding-self/STATEMENT.md)** — for each index, return the product of
  every *other* element, without using division.
  → [products-of-array-discluding-self/submission-3.go, lines 10–14](Data%20Structures%20&%20Algorithms/products-of-array-discluding-self/submission-3.go#L10-L14)
  is the forward pass, [lines 16–20](Data%20Structures%20&%20Algorithms/products-of-array-discluding-self/submission-3.go#L16-L20)
  the backward pass.

### 6. Count what your loop actually iterates

Complexity is measured over what the loop touches, not over `n` by
default. A loop from `min(nums)` to `max(nums)` costs O(range of values).
With values up to a billion that is pseudo-polynomial and times out, even
though the input has only 100k numbers. (The range-scan attempt lived in
chat only, so there is no file to link.)

The second half: when a walk's cost proof depends on visiting each value
once, iterate the deduplicated set, not the raw input. Duplicates in the
input would re-run the walk and break the O(n) argument.

Seen in:
- **[Longest Consecutive Sequence](Data%20Structures%20&%20Algorithms/longest-consecutive-sequence/STATEMENT.md)** — find the length of the longest run of
  consecutive integers hiding in an unsorted list.
  → [longest-consecutive-sequence/submission-9.go, line 22](Data%20Structures%20&%20Algorithms/longest-consecutive-sequence/submission-9.go#L22)
  ranges over the set, not `nums`. [Line 23](Data%20Structures%20&%20Algorithms/longest-consecutive-sequence/submission-9.go#L23)
  is the start-of-sequence gate, [lines 26–28](Data%20Structures%20&%20Algorithms/longest-consecutive-sequence/submission-9.go#L26-L28)
  the upward walk.

### 16. Normalize each element as you read it, not the whole input up front

`strings.ToLower(s)` builds a second string — O(n) extra space spent before
the real work starts. But indexing a string hands you a *copy* of one byte,
and the copy is yours to change even though the string is immutable. If the
algorithm only ever looks at one or two elements at a time, transform them
at read time and no transformed copy of the input ever needs to exist.

Seen in:
- **[Valid Palindrome](Data%20Structures%20&%20Algorithms/is-palindrome/STATEMENT.md)** — ignoring case
  and punctuation, does the string read the same both ways?
  → [is-palindrome/submission-13.go, lines 14 and 23–28](Data%20Structures%20&%20Algorithms/is-palindrome/submission-13.go#L14):
  lowercasing applied to the two bytes being compared. Contrast
  [submission-3.go, lines 5–14](Data%20Structures%20&%20Algorithms/is-palindrome/submission-3.go#L5-L14),
  which built a full cleaned copy first.

Related background: [go-notes/strings.md](go-notes/strings.md).

## Go semantics

### 7. Comma-ok when zero could be real data, direct indexing when it cannot

A map read on a missing key returns the zero value. Direct `if m[k]` is
safe only when that zero value could never be legitimate stored data.
True for bool sets: `false` always means absent. For indices and counts,
0 is real data, so use `value, ok := m[k]`.

Seen in:
- **[Two Sum](Data%20Structures%20&%20Algorithms/two-integer-sum/STATEMENT.md)** —
  [two-integer-sum/submission-4.go, line 4](Data%20Structures%20&%20Algorithms/two-integer-sum/submission-4.go#L4):
  comma-ok, because a stored index of 0 is legitimate.
- **[Longest Consecutive Sequence](Data%20Structures%20&%20Algorithms/longest-consecutive-sequence/STATEMENT.md)** —
  [longest-consecutive-sequence/submission-9.go, lines 23 and 26](Data%20Structures%20&%20Algorithms/longest-consecutive-sequence/submission-9.go#L23-L26):
  direct indexing on a bool set, which is safe.

### 8. The Go set idiom is `map[T]struct{}`

`struct{}` occupies zero bytes and tells the reader that membership is the
only thing stored. `map[T]bool` works, but invites the question of what a
stored `false` would mean.

Where to apply it: the set in
[longest-consecutive-sequence/submission-9.go, line 4](Data%20Structures%20&%20Algorithms/longest-consecutive-sequence/submission-9.go#L4)
is `map[int]bool`. Not wrong, just less idiomatic. Note the trade: with
`struct{}` you lose direct `if set[x]` indexing (takeaway 7) and use
`_, ok :=` instead.

### 9. What can and cannot be a Go map key

Arrays, and structs whose fields are all comparable, compare by content —
so they hash by content and work as keys. Pointers compare by identity:
two pointers to equal content are different keys, and mutating the
pointee does not move the key. Slices, maps, and functions cannot be keys
at all, and a single slice field poisons an otherwise comparable struct.

Seen in:
- **[Group Anagrams](Data%20Structures%20&%20Algorithms/anagram-groups/STATEMENT.md)** —
  [anagram-groups/submission-4.go, line 3](Data%20Structures%20&%20Algorithms/anagram-groups/submission-4.go#L3):
  `map[[26]int][]string`. The array's *content* is the key, which is the
  entire trick — anagrams produce the same count array, so they land in
  the same bucket with no sorting.

### 10. `var` for slices, `make` for maps

A nil slice supports `append`, `len`, and `range`, so `var output []int`
is enough. A nil map allows reads (you get the zero value back) but
panics on write, so a map you intend to fill needs `make`. Pass a
capacity hint only when you actually know one.

Seen in:
- **[Top K Frequent](Data%20Structures%20&%20Algorithms/top-k-elements-in-list/STATEMENT.md)** —
  [top-k-elements-in-list/submission-5.go, lines 7–9](Data%20Structures%20&%20Algorithms/top-k-elements-in-list/submission-5.go#L7-L9):
  `var` for the output slice, `make` for the map and the pre-sized
  buckets, side by side.

### 11. `s[i]` is a byte, `range s` yields runes

Indexing a string gives you raw bytes. Ranging over it decodes UTF-8 and
gives you runes. Byte indexing is correct only when the input is
guaranteed ASCII — say that constraint out loud whenever you lean on it.

Seen in:
- **[Valid Anagram](Data%20Structures%20&%20Algorithms/is-anagram/STATEMENT.md)** —
  [is-anagram/submission-9.go, lines 7–8](Data%20Structures%20&%20Algorithms/is-anagram/submission-9.go#L7-L8):
  byte indexing, justified because inputs are lowercase ASCII.
- **[Encode and Decode Strings](Data%20Structures%20&%20Algorithms/string-encode-and-decode/STATEMENT.md)** — pack a list of strings into one string
  and recover the list exactly.
  → [string-encode-and-decode/submission-13.go, line 24](Data%20Structures%20&%20Algorithms/string-encode-and-decode/submission-13.go#L24):
  `encoded[j] - '0'` reads one digit byte at a known position, which is
  the whole point of the length-prefix format.

### 12. In hot loops, hunt allocations, not function calls

Go function calls cost nanoseconds and small ones inline away. The slow
Group Anagrams attempt was not slow because it called an `isAnagram`
helper. It was slow because it compared all pairs — O(n²·k) — and early
versions allocated a fresh map per pair. When the judge says slow, look
for what the loop allocates or repeats, not for how many functions it
calls.

Seen in:
- **[Group Anagrams](Data%20Structures%20&%20Algorithms/anagram-groups/STATEMENT.md)** — the pairwise shape:
  [anagram-groups/submission-3.go, lines 10–27](Data%20Structures%20&%20Algorithms/anagram-groups/submission-3.go#L10-L27),
  a full inner scan with per-pair counting. The fix,
  [submission-4.go, lines 4–10](Data%20Structures%20&%20Algorithms/anagram-groups/submission-4.go#L4-L10),
  counts each string once and lets the map do the grouping.

## Process

### 13. Constraints first, then ask "what breaks this?"

Write the problem's constraints as comments at the top of the editor
before writing code, then ask what would break the plan before running
it. The pattern of attempt 1 dying on an edge case stopped once this
ritual started.

Seen in:
- [top-k-elements-in-list/submission-5.go, lines 2–6](Data%20Structures%20&%20Algorithms/top-k-elements-in-list/submission-5.go#L2-L6)
- [products-of-array-discluding-self/submission-3.go, lines 2–8](Data%20Structures%20&%20Algorithms/products-of-array-discluding-self/submission-3.go#L2-L8)
- [string-encode-and-decode/submission-13.go, lines 2–6](Data%20Structures%20&%20Algorithms/string-encode-and-decode/submission-13.go#L2-L6)

### 14. Record what attempt 1 died of

The first failure is your signature — it shows which edge case your
instinct misses. Each problem's NOTES.md keeps a "failed attempts"
section for exactly this. Example:
[is-anagram/NOTES.md](Data%20Structures%20&%20Algorithms/is-anagram/NOTES.md).

### 15. Write locally with gofmt-on-save, then clean before pasting

Format on save, delete dead code, and grep for `Println` before every
paste to the judge (the full checklist is in [PLAYBOOK.md](PLAYBOOK.md)).
The standing counterexample:
[anagram-groups/submission-4.go, lines 14–42](Data%20Structures%20&%20Algorithms/anagram-groups/submission-4.go#L14-L42)
carries 29 lines of commented-out old attempt that went to the judge.

### 17. An accepted submission is not a correct program

The judge runs a finite test set. Submission-11 of Valid Palindrome was
accepted while returning `true` for `"@@@@axya"`, whose cleaned form
`"axya"` is not a palindrome — the tests simply never front-load junk
characters. A green tick means "passed these cases", nothing more.

After acceptance, attack your own boundary logic with inputs the tests
might not have: everything-is-skippable, junk skewed to one side, empty
after cleaning, single element. Skip-and-walk code earns this scrutiny
every time.

Seen in:
- **[Valid Palindrome](Data%20Structures%20&%20Algorithms/is-palindrome/STATEMENT.md)** — the accepted
  bug was the loop bound at
  [is-palindrome/submission-11.go, line 8](Data%20Structures%20&%20Algorithms/is-palindrome/submission-11.go#L8);
  the story is in [the NOTES](Data%20Structures%20&%20Algorithms/is-palindrome/NOTES.md).
