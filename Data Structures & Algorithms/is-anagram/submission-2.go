func isAnagram(s string, t string) bool {
	sizes_s := make(map[rune]int)
	sizes_t := make(map[rune]int)
	for _, value := range s {
		sizes_s[value]++
	}
	for _, value2 := range t {
		sizes_t[value2]++
	}
	for _, value := range s {
		if sizes_s[value] != sizes_t[value] {
			return false
		}
	}
	for _, value := range t {
		if sizes_s[value] != sizes_t[value] {
			return false
		}
	}

	return true
}
