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
	// if it starts a sequence i (a number stars a sequence if the number-1 isn't in the set) need to keep track of the longest sequence
	// once i reach the start of a sequence, i just start a loop of item+k
	// to find the longest sequence in the set. 
	// if i find a number in the set that doesn't start a sequence i just 	    // ignore it because i'll get to it anyway in my k loop.
	for item, _ := range set {
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
