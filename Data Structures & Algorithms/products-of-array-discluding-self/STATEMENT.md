# Product of Array Except Self

Given an integer array `nums`, return an array `answer` where `answer[i]`
is the product of every element of `nums` except `nums[i]`. Division is
not allowed, and the target is O(n) time.

**Example**

```
Input:  nums = [1, 2, 4, 6]
Output: [48, 24, 12, 8]
```

**Constraints**

- Array length up to ~10^5.
- Small values (roughly -30 to 30), so products fit in an int.
- Zeroes can appear — the two-pass solution handles them with no special
  cases, unlike the division approach.

**Target:** O(n) time, O(1) extra space — a forward running product and a
backward running product written into the output array.
