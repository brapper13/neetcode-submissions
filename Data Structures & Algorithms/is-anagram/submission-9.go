func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var counts[26] int
	for i := range s {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}
	for _, value := range counts {
		if value != 0 {
			return false
		}
	}
	return true
}
