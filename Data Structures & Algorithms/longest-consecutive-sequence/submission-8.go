func longestConsecutive(nums []int) int {
	// 100k length max
	// numbers large
	set := make(map[int]bool)
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return 1
	}
	largestSeq := 0
	for _, item := range nums {
		set[item] = true
	}
	
	// let's do one pass
	// if i see a number, it can either start a sequence or extend a sequence.
	// if it starts a sequence i need to keep track of the longest sequence
	// to keep track of which sequence the number is tied to we need a map
	// 1, 1, 3, 4, 2, 5
	// 1, 2, 3, 19, 4, 21, 20
	for _, item := range nums {
		if !set[item-1] {
			// start loop
			seq := 0
			for set[item+seq] {
				seq++
			}
			if seq > largestSeq {
				largestSeq = seq
			}
		}
	}
	return largestSeq
}
