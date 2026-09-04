func checkInclusion(s1 string, s2 string) bool {
	// s1 perm meaning needs to contain all letters in s1

	var countS1 [26]int
	var local [26]int
	for _, char := range s1 {
		countS1[char-'a']++
	}
	left := 0
	for i := 0; i < len(s2); i++ {
		local[s2[i]-'a']++
		if i - left + 1 >= len(s1) {
			if countS1 == local {
				return true
			} 
			local[s2[left]-'a']--
			left++
		}
	}
	return false
}
