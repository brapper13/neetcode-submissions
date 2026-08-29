package main

import "testing"

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
	}{
		{[]int{3, 4, 5, 6}, 7},
		{[]int{4, 5, 6}, 10},
		{[]int{5, 5}, 10},
		{[]int{-1, -2, -3, -4, -5}, -8},
		{[]int{0, 4, 3, 0}, 0},
	}
	for _, c := range cases {
		got := twoSum(c.nums, c.target)
		if len(got) != 2 {
			t.Errorf("twoSum(%v, %d) = %v, want two indices", c.nums, c.target, got)
			continue
		}
		i, j := got[0], got[1]
		if i == j || i < 0 || j < 0 || i >= len(c.nums) || j >= len(c.nums) {
			t.Errorf("twoSum(%v, %d) = %v: bad indices", c.nums, c.target, got)
			continue
		}
		if c.nums[i]+c.nums[j] != c.target {
			t.Errorf("twoSum(%v, %d) = %v: values sum to %d", c.nums, c.target, got, c.nums[i]+c.nums[j])
		}
	}
}
