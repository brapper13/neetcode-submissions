# Two Sum II — Input Array Is Sorted

Given a 1-indexed array of integers sorted in non-decreasing order and a
target, return the 1-based indices `[index1, index2]` (index1 < index2)
of the two numbers that add up to the target. Exactly one solution
exists, and you may not use the same element twice.

**Example**

```
Input:  numbers = [1, 2, 3, 4], target = 3
Output: [1, 2]        (1 + 2 = 3)
```

**Constraints**

- Array is sorted, non-decreasing.
- Exactly one solution is guaranteed.
- O(1) additional space required.

**Target:** O(n) time, O(1) space — converging two pointers. Sum too
big moves the right pointer left, too small moves the left pointer
right. The sort order is what makes each move safe to rule out.
