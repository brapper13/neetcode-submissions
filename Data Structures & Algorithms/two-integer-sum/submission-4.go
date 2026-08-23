func twoSum(nums []int, target int) []int {
	sums := make(map[int]int)
    for i, first := range nums {
		if j, ok := sums[target-first]; ok { 
			return []int{j, i}
		}
		sums[first] = i
	}
	return nil
}
