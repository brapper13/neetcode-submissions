func isAnagram(s string, t string) bool {
	sizes := make(map[byte]int)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		sizes[s[i]]++
		sizes[t[i]]--
	}
	for _, value := range sizes {
		if value != 0 {
			return false
		}
	}
	return true
}
