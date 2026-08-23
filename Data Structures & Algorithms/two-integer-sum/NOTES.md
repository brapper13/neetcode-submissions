# Review notes — two-integer-sum (submissions 2–3)

**Verdict:** sub-2 brute force O(n²) correct; sub-3 hash map O(n) correct.
Attempt 1 (failed, unsynced) almost certainly the self-pair trap — [3,2,4],
target 6 → [0,0] — given the `i != j` / `value != j` guards both survivors
carry.

## The canonical one-pass — all the guards evaporate

```go
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int, len(nums))
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return []int{j, i}
		}
		seen[n] = i
	}
	return nil
}
```

**Check before insert.** The map only ever holds elements *earlier* than `i`,
so by construction:
- self-pair impossible (your complement can't be you — you're not inserted yet
  when you check), no `!= j` guard needed;
- `j < i` always, so no ordering swap needed;
- one pass instead of two.

Pattern to pocket: **when guards exist to fix cases the data flow could have
made unrepresentable, look for the shape that makes them impossible.** Here,
insertion order IS the ordering logic.

## Smaller notes

- "Not found" → return `nil`, not `[]int{}` (idiomatic; unreachable here
  anyway — the problem guarantees a solution).
- Problem statement allows indices in any order — the `i <= j` swap was armour
  against a requirement that doesn't exist. Constraints are free points.
- sub-2 mixes tabs and spaces — write locally with gofmt-on-save, paste into
  NeetCode after.
- `sums` holds complement→index; `seen` (element→index, in the one-pass) says
  what it means.

## Milestone

Warm-start set complete: Contains Duplicate, Valid Anagram, Two Sum — all in
Go, all on 2026-08-16, a day before week 2 officially starts.
