func maxProfit(prices []int) int {
	// to make max profit you have to buy low and sell high.
	// if we do a single pass, we can keep track of the lowest and calculate 
    // profit per day. and return highest

	lowest := prices[0]
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] <= lowest {
			lowest = prices[i]
		}
		profit := prices[i] - lowest
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit

}
