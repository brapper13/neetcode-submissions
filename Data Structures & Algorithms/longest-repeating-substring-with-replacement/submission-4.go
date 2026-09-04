func characterReplacement(s string, k int) int {
	// what we can do is keep a window of repeating characters
	// the repeating characters can contain AT MOST k different characters.
	// keep index of current left index of substring.
	// uppercase english characters - can use arrray!
	var count [26]int
	maxFreq := 0
	maxSeq := 0
	left := 0
	for idx, char := range s { 
		count[char - 'A']++
		if count[char - 'A'] > maxFreq {
			maxFreq = count[char - 'A']
		} 
		if (idx - left + 1) - maxFreq > k {
			count[s[left] - 'A']--
			left++
		}
		if (idx - left + 1) > maxSeq {
			maxSeq = idx - left + 1
		}
	}
	return maxSeq
}
