# Top K Frequent Elements

Given an integer array `nums` and an integer `k`, return the `k` values
that occur most often. The answer is guaranteed unique and can be in any
order.

**Example**

```
Input:  nums = [1, 2, 2, 3, 3, 3], k = 2
Output: [2, 3]
```

**Constraints**

- Values between -1000 and 1000 (from the problem page, recorded in the
  solution's header comments).
- The answer is guaranteed to be unique.

**Target:** O(n) via frequency buckets — a frequency can never exceed
`len(nums)`, so an array of `len(nums)+1` buckets replaces sorting.

### Result
Completed. Took 15:16:73. 
I actually had the right answer but due to a silly go mistake spent a lot of time debugging. did for item := range array instead of for _, item
