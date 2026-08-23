func isPalindrome(s string) bool {
	//sanitize the input first
	// s is ascii but case insensitive
	// need to remove non-alphanumeric characters
	s = strings.ToLower(s)
	i := 0
	j := len(s)-1
	for i < len(s)/2 {
		// if front is aln
		for !('a' <= s[i] && s[i] <= 'z' || '0' <= s[i] && s[i] <= '9') && i < j {
			i++
		}
		for !('a' <= s[j] && s[j] <= 'z'|| '0' <= s[j] && s[j] <= '9') && j > i {
			j--
		}

		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
