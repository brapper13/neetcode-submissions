# Best Time to Buy and Sell Stock

Given an array `prices` where `prices[i]` is the price on day `i`,
choose one day to buy and a later day to sell. Return the maximum
profit possible, or 0 if no profitable trade exists.

**Example**

```
Input:  prices = [7, 1, 5, 3, 6, 4]
Output: 5        (buy at 1, sell at 6)
```

**Constraints**

- At least one price; prices non-negative.
- Sell day must be after the buy day.

**Target:** O(n) time, O(1) space — one pass carrying the minimum
price seen so far. Each day, profit-if-sold-today is
`price - minSoFar`; keep the best.
