func twoSum(nums []int, target int) []int {
    for i, first := range nums {
		for j, second := range nums {
			if first + second == target && i!=j {
				if i <= j {
					return []int{i, j}
				}
				return []int{j, i}
			}
		}
	}
	return []int{}
}
