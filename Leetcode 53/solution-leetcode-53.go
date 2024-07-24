func maxSubArray(nums []int) int {
    if len(nums) == 0 {
		return 0
	}
	
	a := nums[0]
	b := nums[0]

	for i := 1; i < len(nums); i++ {
		a = max(nums[i], a + nums[i])
		if a > b {
			b = a
		}
	}

	return b
}