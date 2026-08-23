# Review notes — anagram-groups (submissions 3–4)

**Verdict:** sub-3 pairwise O(n²·k), accepted; sub-4 count-array bucketing
O(n·k) — the canonical answer, and `map[[26]int][]string` is the best key
choice (beats sorted-string keys by skipping the O(k log k) sort).

## The real story of the slow first attempt

The isAnagram-helper version wasn't slow because "Go is bad at function
calls" — calls cost nanoseconds and small functions get inlined. It was slow
because that helper allocates a `map[byte]int` **per call**, called **per
pair**: O(n²) allocations + per-char hashing. Accepted sub-3 swapped in
`var count [26]int` — stack array, zero alloc, direct indexing. Same call
structure would've been fine with the array body.

**Pocket rule: hunt allocations inside hot loops, not function calls around
them.**

## Why `[26]int` works as a map key and `[]int` wouldn't

Map keys need `==` and a stable hash. Arrays are values — the map copies the
key on insert, nothing can alias the copy, size is part of the type. Slices
are headers over shared mutable memory: contents-vs-identity equality is
ambiguous, and mutation through an alias would corrupt the stored key's
hash. Python's tuple-vs-list rule, aliasing edition.

## Style

- `continue InnerLoop` labeled continue in sub-3: nice, correctly used, most
  candidates don't know it exists.
- sub-4 shipped with the whole old solution commented out. Delete dead code
  before submitting — in an interview it reads as clutter and doubt.
- `var output [][]string` over `make([][]string, 0)` (nil slice appends fine).
