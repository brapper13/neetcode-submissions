func threeSum(nums []int) [][]int {
	// no duplicate triplets.
	// naive answer would be to do a triple loop. abhorrent
	// x + y + z = 0; x + y = -z
	// Can we just run 2 sum len(nums) number of times? Not sorted
	var output [][]int
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		item := nums[i]
		// 2 pointer approach
		// let's think about this, the top array is going over each element
		// there's no point having j start from zero because we have already
		// caught any 3 sums possible before i.
		j := i + 1
		k := len(nums) - 1
		target := -item
		// since the input is sorted, we can just skip them over.
		for j < k {
			if nums[j] + nums[k] == target {
				output = append(output, []int{item, nums[j], nums[k]})
				for j < len(nums)-1 && nums[j] == nums[j+1] {
					j++
				}
				for k > 0 && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
				continue
			}
			if nums[j] + nums[k] < target {
				j++
				continue
			}
			k--
		}
		for i < len(nums)-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return output
}
