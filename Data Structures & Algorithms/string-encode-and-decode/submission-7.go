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
		output += fmt.Sprintf("#%s%s", digits, length)
		output += str
	}
	fmt.Println(output)
	return output
}

func (s *Solution) Decode(encoded string) []string {
	var output []string
	var i int
	for i < len(encoded) {
		j := i + 1
		// get digits
		digits, _ := strconv.Atoi(fmt.Sprintf("%c", encoded[j]))
		j++
		var nextLength strings.Builder
		// get length of string
		for k := j + digits; j < k; j++ {
			if _, err := strconv.Atoi(fmt.Sprintf("%c", encoded[j])); err != nil {
				break
			}
			nextLength.WriteByte(encoded[j])
		}
		length, err := strconv.Atoi(nextLength.String())
		if err != nil {
			return output
		}
		var decodedString strings.Builder
		var k int
		for k = 0; k < length; k++ {
			decodedString.WriteByte(encoded[j+k])
		}
		output = append(output, decodedString.String())
		i = j + k
	}
	return output
}
