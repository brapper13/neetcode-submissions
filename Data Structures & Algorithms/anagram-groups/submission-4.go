func groupAnagrams(strs []string) [][]string {
    output := make([][]string, 0)
    seen := make(map[[26]int][]string)
    for _, str := range strs {
		var count [26]int
		for _, char := range str {
			count[char - 'a']++
		}
		seen[count] = append(seen[count], str)
	}
	for _, v := range seen {
		output = append(output, v)
	}
        // if seen[outer] {
        //     continue
        // }
        // sub := make([]string, 0)
		// InnerLoop:
        // for j, inner := range strs {
		// 	// isAnagram := 0
        //     if i == j {
        //         continue
        //     }
		// 	if len(outer) != len(inner) {
        // 		continue
    	// 	}
		// 	var count [26]int
		// 	for k := 0; k < len(inner); k++ {
		// 		count[outer[k]-'a']++
		// 		count[inner[k]-'a']--
		// 	}
		// 	for _, l := range count {
		// 		if l != 0 {
		// 			continue InnerLoop
		// 		}
		// 	}
        //     seen[inner] = true
        //     sub = append(sub, inner)
        // }
        // seen[outer] = true
        // sub = append(sub, outer)
        // output = append(output, sub)
	return output
}