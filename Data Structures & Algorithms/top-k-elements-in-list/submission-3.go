func topKFrequent(nums []int, k int) []int {
	// unique answer
	// any order
	// number between -1000 to 1000
	// what happens when k is 0?
	var output []int
	counts := make(map[int]int) // map of number to frequency
	values := make(map[int][]int) // map of frequency to list of number
	var sortedValues []int
	for _, value := range nums {
		counts[value]++
	}

	for key, value := range counts {
		values[value] = append(values[value], key)
	}

	for freq := range values {
		sortedValues = append(sortedValues, freq)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(sortedValues)))

	for _, sorted := range sortedValues {
		for _, num := range values[sorted] {
			output = append(output, num)
			if len(output) == k {
				return output
			}
		}
	}
	return output
}
