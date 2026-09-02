# Trapping Rain Water

Given an array `height` where `height[i]` is the height of a bar of
width 1, compute how much water the bars trap after raining.

**Example**

```
Input:  height = [0, 2, 0, 3, 1, 0, 1, 3, 2, 1]
Output: 9
```

**Constraints**

- Array length up to ~2 * 10^4; heights non-negative.
- At least one bar is guaranteed on LeetCode; an empty array would
  need a guard this solution doesn't carry.

**Target:** O(n) time. Per position: water is
`min(maxLeft, maxRight) - height[k]` with both maxes *inclusive* of
`k`, which makes negatives impossible. Two prefix/suffix passes give
O(n) space; the two-pointer refinement reaches O(1) space.
