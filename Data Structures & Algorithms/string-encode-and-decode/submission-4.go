type Solution struct{}
// we're using ascii characters, 256 of them.
// string can be 200 chars
// array can be empty
// array can have 100 strings.
// if adding all the strings up, how to know where is the split?
func (s *Solution) Encode(strs []string) string {
	var output string
	if len(strs) == 0 {
		return "-100"
	}
	for idx, s := range strs {
		output += s
		if idx != len(strs) - 1 {
			output += "#!~/100"
		}	
	}
	return output
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "-100" {
		return []string{}
	}
	input := strings.Split(encoded, "#!~/100")
	fmt.Println(input)
	return input
}
