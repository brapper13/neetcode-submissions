# Group Anagrams

Given an array of strings `strs`, group the anagrams together and return
the groups. Groups and their contents can come in any order.

**Example**

```
Input:  strs = ["act", "pots", "tops", "cat", "stop", "hat"]
Output: [["hat"], ["act", "cat"], ["stop", "pots", "tops"]]
```

**Constraints**

- Up to ~10^4 strings, each up to ~100 chars.
- Lowercase English letters only.

**Target:** O(n·k) — one `[26]int` count per string, used directly as the
map key. No sorting, no pairwise comparison.
