func trap(height []int) int {
	// interesting problem. The amount of water that can be trapped between
	// bars is - 
	// a. starts from the first bar.
	// b. from the end, it starts from the first descending bar
	// okay another solution,
	// water on top of level k, is min(maxLeft, maxRight) - height(k)
	// how to calculate maxLeft and maxRight?
	maxLeft := make([]int, len(height))
	maxLeft[0] = height[0]
	for idx := 1; idx < len(height); idx++ {
		if height[idx] > maxLeft[idx-1] {
			maxLeft[idx] = height[idx]
			continue
		}
		maxLeft[idx] = maxLeft[idx-1]
	}
	maxRight := make([]int, len(height))
	maxRight[len(height)-1] = height[len(height)-1]
	for idx := len(height)-2; idx >= 0; idx-- {
		if height[idx] > maxRight[idx+1] {
			maxRight[idx] = height[idx]
			continue
		}
		maxRight[idx] = maxRight[idx+1]
	}

	totalWater := 0
	for idx := 0; idx < len(height); idx++ {
		// water level at k = min(maxLeft, maxRight) - height(k)
		totalWater += min(maxLeft[idx], maxRight[idx]) - height[idx]
	}
	return totalWater


}
