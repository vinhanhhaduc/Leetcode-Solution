func canSortArray(nums []int) bool {
	maxa := len(nums) - 1
	noa := 0
	for noa < maxa {
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				if helper(nums[i]) == helper(nums[i+1]) {
					nums[i] += nums[i+1]
					nums[i+1] = nums[i] - nums[i+1]
					nums[i] = nums[i] - nums[i+1]
					noa = 0
				} else {
					return false
				}
			} else {
				noa++
			}
		}
	}
	return true
}

func helper(num int) int {
	count := 0
	for num > 0 {
		count += num & 1
		num >>= 1 
	}
	return count
}