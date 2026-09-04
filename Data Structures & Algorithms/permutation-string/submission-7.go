func checkInclusion(s1 string, s2 string) bool {
	// s1 perm meaning needs to contain all letters in s1

	// change to map
	var countS1 [26]int
	var local [26]int
	for _, char := range s1 {
		countS1[char-'a']++
	}

	Outer: for i := 0; i < len(s2) - len(s1) + 1; i++ {
		local = countS1
		for j := 0; j < len(s1); j++ {
			local[s2[i+j] - 'a']--
		}
		for _, item := range local {
			if item != 0 {
				continue Outer
			}
		}
		return true
	}
	return false
}
