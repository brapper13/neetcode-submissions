func isPalindrome(s string) bool {
	//sanitize the input first
	// s is ascii but case insensitive
	// need to remove non-alphanumeric characters
	s = strings.ToLower(s)
	var newS strings.Builder
	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z' || 'A' <= b && b <= 'Z' || '0' <= b && b <= '9') {

			newS.WriteByte(b)
		}
	}
	s = newS.String()
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
