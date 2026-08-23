# Review notes — is-anagram (submissions 2–7)

**Verdict:** correct from sub-2 onward; the progression is the story.
2 (two maps, four loops) → 4 (length guard + single map ±1, the canonical
answer) → 5 (bytes over runes) → 7 (guard before allocating). Reaching
canonical form within one sitting is the practice working.

## Final version (sub-7), annotated

- `s[i]` is a **byte**; `range s` would yield **runes** (decoded UTF-8).
  Byte indexing is correct here only because inputs are guaranteed lowercase
  ASCII — say that constraint out loud in an interview. Unicode-safe version
  uses `map[rune]int` + `range`. (Drill F09 covers this.)
- Guarding `len(s) != len(t)` before allocating the map: good instinct.
  Note `len(s)` is *byte* length either way.

## Upgrade unlocked by the constraint

Bounded alphabet → fixed array beats a map (no hashing, no allocation):

```go
func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var counts [26]int
	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}
	for _, c := range counts {
		if c != 0 {
			return false
		}
	}
	return true
}
```

Pattern to pocket: **when the alphabet is bounded, use an array indexed by
`ch - 'a'` instead of a map.** Recurs in sliding-window and bitmask problems.

## Idiom

- `sizes_s` → Go locals are lowerCamel (`countS`), no snake_case.
- Two params of one type: `func isAnagram(s, t string)`.

## Failed attempts

1 and 6 were wrong answers (test cases failed), not compile errors — the Go
compiled; the logic didn't survive the judge. Specific breaking case not
recalled; likely candidates given the surviving code: unequal lengths, or a
char present in `t` only (the case one-directional count comparison misses).
