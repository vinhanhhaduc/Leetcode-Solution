func longestSubarray(nums []int) int {
    maxVal := nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
		}
	}
	maxLen := 0
	currentLen := 0
	for i := 0; i < n; i++ {
		if nums[i] == maxVal {
			currentLen++
			if currentLen > maxLen {
				maxLen = currentLen
			}
		} else {
			currentLen = 0
		}
	}

	return maxLen
}