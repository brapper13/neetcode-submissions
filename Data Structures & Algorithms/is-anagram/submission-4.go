func isAnagram(s string, t string) bool {
	sizes_s := make(map[rune]int)
	if len(s) != len(t) {
		return false
	}
	for _, value := range s {
		sizes_s[value]++
	}
	for _, value2 := range t {
		sizes_s[value2]--
	}
	for _, value := range sizes_s {
		if value != 0 {
			return false
		}
	}
	return true
}
