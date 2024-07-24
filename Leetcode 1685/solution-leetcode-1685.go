func getSumAbsoluteDifferences(nums []int) []int {
    n := len(nums)
	res := make([]int, n)

	prefSum := 0
	suffSum := 0
	for _, num := range nums {
		suffSum += num
	}

	for i := 0; i < n; i++ {
		suffSum -= nums[i]
		if i > 0 {
			prefSum += nums[i-1]
		}
		l := nums[i] * i - prefSum
		r := suffSum - nums[i] * (n - i - 1)
		res[i] = l + r
	}

	return res
}