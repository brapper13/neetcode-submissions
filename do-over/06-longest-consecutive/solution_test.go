package main

import "testing"

func TestLongestConsecutive(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{2, 20, 4, 10, 3, 4, 5}, 4},
		{[]int{100, 4, 200, 1, 3, 2}, 4},
		{[]int{0, -1, -2, 5}, 3},
		{[]int{1, 2, 0, 1}, 3},
		{[]int{5, 5, 5}, 1},
	}
	for _, c := range cases {
		if got := longestConsecutive(c.nums); got != c.want {
			t.Errorf("longestConsecutive(%v) = %d, want %d", c.nums, got, c.want)
		}
	}
}
