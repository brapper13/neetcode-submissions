func topKFrequent(nums []int, k int) []int {
	// unique answer
	// any order
	// number between -1000 to 1000
	// what happens when k is 0?
	// the output can only be as large as k. The frequency can only be as 		large as the number of items as well.
	var output []int
	counts := make(map[int]int, len(nums)) // map of number to frequency
	values := make([][]int, len(nums)+1, len(nums)+1) // frequency bucket
	for _, value := range nums {
		counts[value]++
	}

	for key, value := range counts {
		values[value] = append(values[value], key)
	}

	for i := len(values) - 1; i >= 0; i-- {
		for _, num := range values[i] {
			output = append(output, num)
			if len(output) == k {
				return output
			}
		}
	}
	return output
}
