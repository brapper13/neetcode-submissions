func lengthOfLongestSubstring(s string) int {
	left := 0
	seen := make(map[rune]int)
	maxSeq := 0
	for idx, char := range s {
		if _, ok := seen[char]; ok {
			// the max is what lets us handle stale duplicates.
			// a stale duplicate is a duplicate that's behind left so 
			// when encountering in the current sequence it doesn't break the sequence so this keeps left consistent.
			left = max(left, seen[char] + 1)
			seen[char] = idx
		} else {
			seen[char] = idx
		}
		if (idx - left + 1) >= maxSeq {
			maxSeq = idx - left + 1
		}
	}
	return maxSeq
}
