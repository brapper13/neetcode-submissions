# Valid Anagram

Given two strings `s` and `t`, return `true` if `t` is an anagram of `s` —
that is, `t` uses exactly the same characters as `s` with exactly the same
counts, in any order.

**Example**

```
Input:  s = "racecar", t = "carrace"
Output: true

Input:  s = "jar", t = "jam"
Output: false
```

**Constraints**

- Both strings consist of lowercase English letters only (ASCII).
- Length up to ~10^4.

**Target:** O(n) time. The lowercase-only constraint is what allows a
`[26]int` counter instead of a map.

### Result
format - min:sec:msec
29/08/2026: passed. 4:15:06
