package main

func longestConsecutive(nums []int) int {
	// your rewrite: iterate the set, no dead code
	// keyword: consecutive. meaning that once we find the start of a sequence we can just   // know the length of the sequence starting with that number in the list. so on, until   // we find the longest sequence.
	numCheck := make(map[int]bool)
	for _, value := range nums {
		numCheck[value] = true
	}

	var longest int
	for value := range numCheck {
		// start of the sequence
		if !numCheck[value-1] {
			seqLength := 1
			for i := value + 1; ; i++ {
				if !numCheck[i] {
					if seqLength >= longest {
						longest = seqLength
					}
					break
				}
				seqLength++
			}
		}
	}

	return longest
}
