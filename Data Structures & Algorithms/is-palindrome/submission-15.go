func isPalindrome(s string) bool {
	// s is ascii but case insensitive
	// need to ignore non-alphanumeric characters
	i := 0
	j := len(s) - 1
	for i < j {
		for i < j && !isAlphaNumeric(s[i]) {
			i++
		}
		for j > i && !isAlphaNumeric(s[j]) {
			j--
		}

		if toLowerCase(s[i]) != toLowerCase(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

func toLowerCase(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		return b + 'a' - 'A'
	}
	return b
}

func isAlphaNumeric(b byte) bool {
	return'a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || '0' <= b && b <= '9'
}
