func isAnagram(s string, t string) bool {
	sizes_s := make(map[rune]int)
	sizes_t := make(map[rune]int)
	for _, value := range s {
		sizes_s[value]++
		sizes_t[value]--
	}
	for _, value2 := range t {
		sizes_t[value2]++
		sizes_s[value2]--
	}
	for _, value := range sizes_s {
		if value > 0 {
			return false
		}
	}
	for _, value := range sizes_t {
		if value > 0 {
			return false
		}
	}

	return true
}
