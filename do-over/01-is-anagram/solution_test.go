package main

import "testing"

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		s, t string
		want bool
	}{
		{"racecar", "carrace", true},
		{"jar", "jam", false},
		{"a", "ab", false},
		{"", "", true},
		{"aabb", "abab", true},
		{"aacc", "ccac", false},
	}
	for _, c := range cases {
		if got := isAnagram(c.s, c.t); got != c.want {
			t.Errorf("isAnagram(%q, %q) = %v, want %v", c.s, c.t, got, c.want)
		}
	}
}
