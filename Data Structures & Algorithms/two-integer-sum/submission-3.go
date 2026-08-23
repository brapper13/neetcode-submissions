func twoSum(nums []int, target int) []int {
	sums := make(map[int]int)
    for i, first := range nums {
		sums[target-first] = i
	}
	for j, second := range nums {
		if value, ok := sums[second]; ok {
			if value != j {
				if j < value {
					return []int{j, value}
				}
				return []int{value, j}
			}
		}
	}
	return []int{}
}
