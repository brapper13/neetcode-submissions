func productExceptSelf(nums []int) []int {
	// no strings yay!
	// array length 100k
	// -30 <= int <= 30
	// if one zero, all products except zero index is zero
	// if 2 zeroes, all products are zero
	// another solution would be for item idx i in the array, the output[i] is product of everything before it multiplied by the product of everything ahead of it. We can do this with a multi-pass.
	preProducts := make(map[int]int)
	postProducts := make(map[int]int)
	output := make([]int, len(nums), len(nums))
	preProducts[0] = 1
	postProducts[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		preProducts[i] = preProducts[i-1] * nums[i - 1]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		postProducts[i] = postProducts[i+1] * nums[i+1]
	}
	for i, _ := range output {
		output[i] = preProducts[i] * postProducts[i]
	}
	return output
}
