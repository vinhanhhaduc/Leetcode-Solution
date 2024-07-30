
func countSubarrays(nums []int, k int) int64 {
	var res int64 = 0
	var i int64 = 0
	max := 0
	for _, val := range nums {
		if max < val {
			max = val
		}
	}

	count := 0

	for _, val := range nums {
		if val == max {
			count++
		}

		for count >= k {
			if nums[i] == max {
				count--
			}
			i++
		}
		res += i
	}

	return res
}