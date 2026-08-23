# Review notes — top-k-elements-in-list (submissions 2–4)

**Verdict:** sub-4 is the canonical O(n) bucket solution, derived from hints
alone across three iterations. Best problem of the prep so far.

## The arc

- sub-2: map[freq][]num + sort of unique freqs, m/index counter gymnastics,
  latent tie-at-cutoff overflow (masked by the unique-answer guarantee).
- sub-3: renamed shadowing loop var (`freq`), replaced counters with
  `len(output) == k` early return — which deleted the latent bug as a side
  effect. Simpler and safer were the same edit.
- sub-4: derived the bound ("frequency can only be as large as the number of
  items"), swapped map+sort for a bucket slice indexed by frequency, walked
  backwards. Sort deleted, O(n log n) → O(n).

## Pattern to pocket

Same move as anagram-groups sub-3→4: the idea was already present, stored in
the wrong container. **When map keys are small dense bounded ints, a slice
indexed by the key replaces map + sort.** Third appearance of
bounded-domain-beats-hashing ([26]int counts, [26]int key, freq buckets).

## Cosmetic residue

- `make([][]int, n+1, n+1)` → `make([][]int, n+1)` (third arg only when
  cap > len).
- `values` named for its map life; it's `buckets` now.
- buckets[0] is permanently empty (no zero frequencies) — fine, know it.
- Constraint-comment ritual at top: retained, now includes the bound
  derivation. Keep this habit forever.
