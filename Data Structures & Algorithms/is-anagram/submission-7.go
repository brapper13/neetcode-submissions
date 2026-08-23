func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sizes := make(map[byte]int)
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
