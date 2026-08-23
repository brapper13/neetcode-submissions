func productExceptSelf(nums []int) []int {
	// no strings yay!
	// array length 100k
	// -30 <= int <= 30
	// if one zero, all products except zero index is zero
	// if 2 zeroes, all products are zero
	// another solution would be for item idx i in the array, the output[i] is product of everything before it multiplied by the product of everything ahead of it. We can do this with a multi-pass.
	// we can use a single ds instead of multi-ds
	out := make([]int, len(nums))
	left := 1
	for i := 0; i < len(nums); i++ {
		out[i] = left
		left *= nums[i]
	}

	right := 1
	for i := len(nums)-1; i >= 0; i-- {
		out[i] *= right
		right *= nums[i]
	}
	return out

}
