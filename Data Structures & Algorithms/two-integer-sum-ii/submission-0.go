func twoSum(numbers []int, target int) []int {
	// sorted in non-decreasing order
	// 1 index array [index1, index2]
	// index1 < index2 && index1 != index2. One valid solution.
	// O(1) additional space
	// Since the array is sorted, let's try using the 2 pointer algorithm
	i := 0
	j := len(numbers) - 1
	for i < j {
		if numbers[i] + numbers[j] == target {
			return []int{i+1, j+1}
		}
		// smallest + largest > target means that j--
		if numbers[i] + numbers[j] > target {
			j--
			continue
		} else {
			i++
		}
	}
	return []int{}
}


