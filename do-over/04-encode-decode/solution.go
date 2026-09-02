package main

import (
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	// your solution
	// to distinuguish between strings we need a delimiter.
	// we need to add length of string so that we know how many characters to read after the delimiter.
	var output strings.Builder
	for _, str := range strs {
		output.WriteString(strconv.Itoa(len(str)))
		output.WriteByte('#')
		output.WriteString(str)
		// Using string += string in a loop is expensive.
		// Using writeString with sprintf is expensive.
		// delimiter := fmt.Sprintf("%d#", len(str))
		// output += fmt.Sprintf("%s%s", delimiter, str)
		// or
		// output.WriteString(fmt.Sprintf("%s%s", delimiter, str))
	}
	// apparently doing string += string in a loop is inefficient.
	return output.String()
}

func (s *Solution) Decode(encoded string) []string {
	// your solution
	// this has been my bane. I think the best way to appraoch a loop like this where
	// you basically don't care about incrementing i in go is to -
	// 1. The top level loop can just be a basic conditional i.e i < len(x)
	// 2. the sub loops would then basically be conditional to get to the delimiter and they don't update i.
	// 3. Only update i when you know what to update it by.
	// 4. Always remember the start of the index is always inclusive, and the last is exclusive. x[i:i+2] is inclusive i and not includiing i + 2. So it has i and i + 1
	// This means that our i should always be at the start of what we want to extract.
	var output []string
	for i := 0; i < len(encoded); {
		// extract string
		var j int
		for encoded[i+j] != '#' {
			j++
		}
		strLength, _ := strconv.Atoi(encoded[i : i+j])
		i += j + 1
		output = append(output, encoded[i:i+strLength])
		i += strLength
	}
	return output
}
