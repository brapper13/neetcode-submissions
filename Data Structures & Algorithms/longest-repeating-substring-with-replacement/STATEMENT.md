# Longest Repeating Character Replacement

Given a string `s` of uppercase English letters and an integer `k`,
you may replace at most `k` characters. Return the length of the
longest substring that can be made into a single repeated character.

**Example**

```
Input:  s = "AABABBA", k = 1
Output: 4        (replace one to get "ABBB" -> "BBBB" or "AABA" -> "AAAA")
```

**Constraints**

- Uppercase English letters only — a [26]int counts array applies.
- k >= 0.

**Target:** O(n) time — sliding window over counts. A window is
fixable when `length - maxFreq <= k`, where maxFreq is the count of
its most frequent character. Grow right, creep left on violation,
score every iteration.
