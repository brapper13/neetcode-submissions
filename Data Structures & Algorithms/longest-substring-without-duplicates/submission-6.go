func lengthOfLongestSubstring(s string) int {
	left := 0
	seen := make(map[rune]int)
	maxSeq := 0
	currSeq := 0
	for idx, char := range s {
		if _, ok := seen[char]; ok {
			if seen[char] < left {
				seen[char] = idx
				currSeq += 1
				if currSeq > maxSeq {
					maxSeq = currSeq
				}
				continue
			}
			left = max(left, seen[char] + 1)
			currSeq = idx - left + 1
			seen[char] = idx
		} else {
			seen[char] = idx
			currSeq += 1
			if currSeq > maxSeq {
				maxSeq = currSeq
			}
		}
	}
	return maxSeq
}
