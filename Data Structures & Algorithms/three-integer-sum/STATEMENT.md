# 3Sum

Given an integer array `nums`, return all triplets `[nums[i], nums[j],
nums[k]]` such that `i`, `j`, `k` are distinct indices and the three
values sum to zero. The output must not contain duplicate triplets.
Triplet order and order within a triplet do not matter.

**Example**

```
Input:  nums = [-1, 0, 1, 2, -1, -4]
Output: [[-1, -1, 2], [-1, 0, 1]]
```

**Constraints**

- Array length up to ~1000.
- Values roughly in [-10^5, 10^5].

**Target:** O(n²) time after an O(n log n) sort. Sort, then for each
anchor run the Two Sum II converging two-pointer on the remainder.
Dedup comes from the sort: skip equal anchors, and skip equal `j`
values after a hit.
