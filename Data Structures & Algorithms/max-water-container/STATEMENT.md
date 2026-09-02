# Container With Most Water

Given an array `heights` where `heights[i]` is the height of a vertical
line at position `i`, find two lines that together with the x-axis form
a container holding the most water. Return the maximum area. Width is
the index distance, height is the shorter of the two lines.

**Example**

```
Input:  heights = [1, 7, 2, 5, 4, 7, 3, 6]
Output: 36        (between the 7 at index 1 and the 6 at index 7)
```

**Constraints**

- Array length up to ~10^5, so O(n²) pair-checking is out.

**Target:** O(n) time, O(1) space — converging two pointers. Always
move the shorter side inward: the shorter line caps the area, so every
pair it could form with a nearer opposite line is provably no better.
