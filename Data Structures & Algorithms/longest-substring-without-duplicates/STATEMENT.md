# Longest Substring Without Repeating Characters

Given a string `s`, return the length of the longest substring
(consecutive characters) that contains no repeated character.

**Example**

```
Input:  s = "zxyzxbb"
Output: 4        ("yzxb")
```

**Constraints**

- Length up to ~5 * 10^4; printable ASCII.

**Target:** O(n) time — sliding window. Map each character to its
last-seen index. On a duplicate, jump the left edge to just past the
previous occurrence, guarded so the left edge never moves backwards:
`left = max(left, prev+1)`. The guard is what makes stale map entries
harmless.
