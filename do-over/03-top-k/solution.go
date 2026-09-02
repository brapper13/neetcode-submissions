package main

func topKFrequent(nums []int, k int) []int {
	// your solution
	// unique answer
	// values between -1000 and 1000
	// using a frequency bucket, naturally sorted.
	counts := make(map[int]int)
	freqBucket := make([][]int, len(nums)+1)
	var output []int
	// get mapping of num(idx) -> frequency
	for _, value := range nums {
		counts[value]++
	}
	// loop through counts, create frequency bucket
	// mapping of frequency to array of items that have that frequency
	for num, freq := range counts {
		freqBucket[freq] = append(freqBucket[freq], num)
	}

	// rev loop through freq bucket, if
	for i := len(freqBucket) - 1; i >= 0; i-- {
		for _, item := range freqBucket[i] {
			output = append(output, item)
		}
		if len(output) >= k {
			return output[0:k]
		}
	}
	return output
}
