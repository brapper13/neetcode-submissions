# Permutation in String

Given two strings `s1` and `s2`, return true if `s2` contains any
permutation of `s1` as a contiguous substring.

**Example**

```
Input:  s1 = "abc", s2 = "lecabee"
Output: true        ("cab" is a permutation of "abc")
```

**Constraints**

- Lowercase English letters — [26]int counts apply.
- A permutation has exactly len(s1) characters, so only windows of
  that one size matter.

**Target:** O(n) time — fixed-size (lockstep) sliding window. One
count array for s1, one for the window; each step one character
enters, one falls off the back, compare with `==` (arrays are
comparable values).
