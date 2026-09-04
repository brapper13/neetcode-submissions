func characterReplacement(s string, k int) int {
	// uppercase english characters - can use arrray!
	maxSeq := 0
	for i := range 26 {
		bad := 0
		left := 0
		for _, char := range s { 
			if int(char - 'A') != i {
				bad++
			}
			if bad > k {
				if int(s[left]-'A') != i {
					bad--
				}
				left++
			}
		}
		maxSeq = max(maxSeq, len(s) - left)
	}
	return maxSeq
}
