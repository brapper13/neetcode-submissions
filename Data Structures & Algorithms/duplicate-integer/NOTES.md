# Review notes — submission-1.go

**Verdict:** correct. O(n) time, O(n) space, early return on first duplicate.
All notes below are idiom, not correctness — the gap between "translated
Python" and native Go.

## Idiomatic version

```go
func hasDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		if _, ok := seen[n]; ok {
			return true
		}
		seen[n] = struct{}{}
	}
	return false
}
```

## Notes

1. **Set idiom: `map[int]struct{}`, not `map[int]int` with `= 1`.**
   `struct{}` occupies zero bytes, and it signals "membership only, the value
   is meaningless" to every Go reader. A dict-with-dummy-value is the Python
   accent for the same thing.

2. **`for _, n := range nums`, not a C-style indexed loop.**
   Declaring `item := 0` before the loop and assigning inside is a C habit;
   Go scopes the element variable inside the loop and skips the index you
   never use. (Remember `range` hands you a *copy* of the element — fine
   here, a trap when mutating; see drill F02.)

3. **Pre-size the map: `make(map[int]struct{}, len(nums))`.**
   Avoids incremental rehashing as the map grows. Worth one spoken clause in
   an interview, no more.

4. **Naming: `seen` over `m`.** Interview code is narrated aloud;
   self-describing names do part of the narration for you.
