# Longest Consecutive Sequence

Given an unsorted integer array `nums`, return the length of the longest
run of consecutive integers that appears in it. The run's elements can sit
anywhere in the array, in any order. Required complexity is O(n).

**Example**

```
Input:  nums = [2, 20, 4, 10, 3, 4, 5]
Output: 4        (the run 2, 3, 4, 5)
```

**Constraints**

- Array length up to ~10^5.
- Values up to ±10^9 — which is why scanning the value range instead of
  the elements is pseudo-polynomial and times out.

**Target:** O(n) — build a set, iterate the set (not the input), only walk
upward from numbers that start a run (`num-1` absent from the set).

### Complete
Result: 09:21:09
1. Remember to iterate over the set at the end otherwise duplicates becomes an issue
2. if seqLength >= longest, not just greater than.
