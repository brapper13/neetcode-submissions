# Two Sum

Given an array of integers `nums` and an integer `target`, return the
indices of the two numbers that add up to `target`. Exactly one valid pair
exists, and you may not use the same element twice.

**Example**

```
Input:  nums = [3, 4, 5, 6], target = 7
Output: [0, 1]        (3 + 4 = 7)
```

**Constraints**

- Array length up to ~10^4.
- Exactly one solution is guaranteed.

**Target:** O(n) time, one pass — map from value to index, check the
complement before inserting.
