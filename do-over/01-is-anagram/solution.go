package main

func isAnagram(s string, t string) bool {
	// your solution
	// ascii chars only - allows us to use an bounded array [26]
	// otherwise we could've used a map.
	if len(s) != len(t) {
		return false
	}
	var counts [26]int
	for idx := range s {
		counts[s[idx]-'a']++
		counts[t[idx]-'a']--
	}

	for _, v := range counts {
		if v != 0 {
			return false
		}
	}

	return true
}
