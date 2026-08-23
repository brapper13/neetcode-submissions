func productExceptSelf(nums []int) []int {
	// no strings yay!
	// array length 100k
	// -30 <= int <= 30
	// if one zero, all products except zero index is zero
	// if 2 zeroes, all products are zero
	products := make([]int, len(nums), len(nums))
	sum := 1
	numZeros := 0
	for _, num := range nums {
		if num == 0 {
			numZeros++
			continue
		}
		sum *= num
	}
	for i, _ := range products {
		if numZeros > 1 {
			products[i] = 0
			continue
		}
		if numZeros == 1 {
			if nums[i] == 0 {
				products[i] = sum
				continue
			}
			products[i] = 0
			continue
		}
		products[i] = sum/nums[i]
	}
	return products
}
