package main

import (
	"strings"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	cases := [][]string{
		{},
		{""},
		{"", ""},
		{"go", "neet"},
		{"#", "12#4"},
		{"a#3b", "", "xy"},
		{"4neet"},
		{"line\nbreak", "with space"},
		{strings.Repeat("x", 250)},
	}
	s := &Solution{}
	for _, in := range cases {
		got := s.Decode(s.Encode(in))
		if len(got) != len(in) {
			t.Errorf("round trip of %q: got %q", in, got)
			continue
		}
		for i := range in {
			if got[i] != in[i] {
				t.Errorf("round trip of %q: got %q", in, got)
				break
			}
		}
	}
}
