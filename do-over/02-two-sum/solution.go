package main

// There is another way to solve twoSum if the list was sorted, we could use the 2 pointers method
// If not sorted, we do this.
func twoSum(nums []int, target int) []int {
	// your solution
	// one solution guaranteed. large array
	// not really a sumMap - bad name.
	// What we're doing here is keeping a map of numbers->indexes we've already seen.
	sumMap := make(map[int]int)
	for idx, value := range nums {
		// If the key target-current is already in the map, then the result is the current index, result index.
		if _, ok := sumMap[target-value]; ok {
			return []int{idx, sumMap[target-value]}
		}
		sumMap[value] = idx
	}
	return []int{}
}
