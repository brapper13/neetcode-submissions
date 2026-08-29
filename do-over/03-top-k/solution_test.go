package main

import (
	"sort"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want []int // any order
	}{
		{[]int{1, 2, 2, 3, 3, 3}, 2, []int{2, 3}},
		{[]int{7, 7}, 1, []int{7}},
		{[]int{1}, 1, []int{1}},
		{[]int{-1, -1, 0}, 2, []int{-1, 0}},
		{[]int{4, 4, 4, 5, 5, 6}, 2, []int{4, 5}},
	}
	for _, c := range cases {
		got := append([]int(nil), topKFrequent(c.nums, c.k)...)
		want := append([]int(nil), c.want...)
		sort.Ints(got)
		sort.Ints(want)
		if len(got) != len(want) {
			t.Errorf("topKFrequent(%v, %d) = %v, want %v", c.nums, c.k, got, want)
			continue
		}
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("topKFrequent(%v, %d) = %v, want %v", c.nums, c.k, got, want)
				break
			}
		}
	}
}
