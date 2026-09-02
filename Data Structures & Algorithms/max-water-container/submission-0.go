func maxArea(heights []int) int {
	i := 0
	j := len(heights) - 1
	maxArea := 0
	for i < j {
		if (j-i)*min(heights[i], heights[j]) > maxArea {
			maxArea = (j - i) * min(heights[i], heights[j])
		}
		if heights[i] <= heights[j] {
			i++
			continue
		} else {
			j--
		}
	}
	return maxArea

}