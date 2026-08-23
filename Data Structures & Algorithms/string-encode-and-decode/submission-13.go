type Solution struct{}
// we're using ascii characters, 256 of them.
// string can be 200 chars
// array can be empty
// array can have 100 strings.
// if adding all the strings up, how to know where is the split?
func (s *Solution) Encode(strs []string) string {
	var output string
	for _, str := range strs {
		length := fmt.Sprintf("%d", len(str))
		digits := fmt.Sprintf("%d", len(length))
		output += fmt.Sprintf("%s%s", digits, length)
		output += str
	}
	return output
}

func (s *Solution) Decode(encoded string) []string {
	var output []string
	var i int
	for i < len(encoded) {
		j := i
		// get digits
		digits := int(encoded[j] - '0')
		j++
		// get length of string
		// i was using a strings.Builder before.
		length, _ := strconv.Atoi(encoded[j:j+digits])
		j+=digits
		output = append(output, encoded[j:j+length])
		i = j + length
	}
	return output
}
