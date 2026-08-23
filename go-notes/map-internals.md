# Map internals: buckets, growth, pre-sizing

## The classic layout

Go's map (through 1.23) is a hash table of buckets, each holding up to 8
key/value pairs plus a per-slot hash fragment for fast rejection. A full
bucket chains to an overflow bucket.

When the average load passes about 6.5 pairs per bucket, the map grows:
the bucket array doubles, and entries are *not* moved all at once. Each
subsequent insert or delete evacuates a little of the old table —
incremental rehashing, so no single operation pays the whole cost.

Go 1.24 replaced this with a Swiss-table design (grouped control bytes
probed in parallel), but the observable behaviour and the advice below are
unchanged.

## What this means in practice

- **Pre-size when you know the count**: `make(map[K]V, n)` allocates for
  `n` entries up front and skips the grow-evacuate churn. This is why the
  top-k solution passes `len(nums)` to `make`.
- **Every map is an allocation**, plus more as it grows. A map created
  inside a hot loop is the classic hidden cost — it's what actually made
  the pairwise Group Anagrams attempt slow.
- **Iteration order is randomized** on purpose, per iteration, so nobody
  can depend on it. If you need order, collect keys and sort.
- **A nil map** (declared with `var`) supports reads — you get the zero
  value — but panics on write. Maps you intend to fill need `make`.
  Slices differ: a nil slice supports `append`, `len`, and `range`.
