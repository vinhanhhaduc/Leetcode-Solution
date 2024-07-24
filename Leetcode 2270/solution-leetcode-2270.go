func waysToSplitArray(nums []int) int {
    n := len(nums)
	if n < 2 {
		return 0
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	prefSum := 0
	res := 0
	for i := 0; i < n-1; i++ {
		prefSum += nums[i]
		if prefSum >= (sum - prefSum) {
			res++
		}
	}

	return res
}