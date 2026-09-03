func trap(height []int) int {
	// interesting problem.
	// water on top of level k, is min(maxLeft, maxRight) - height(k)
	// how to calculate maxLeft and maxRight?
	i := 0
	j := len(height) - 1
	maxL := height[i]
	maxR := height[j]
	totalWater := 0
	for i < j {
		if height[i] > maxL {
			maxL = height[i]
		}
		if height[j] > maxR {
			maxR = height[j]
		}
		// if maxL is less than maxR, then min(maxL, maxR) will be maxL
		if maxL <= maxR {
			totalWater += maxL - height[i]
			i++
		}
		if maxL > maxR {
			totalWater += maxR - height[j]
			j--
		}
	}

	return totalWater


}
