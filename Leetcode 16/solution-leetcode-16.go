func threeSumClosest(nums []int, target int) int {
    sort.Ints(nums)
	n = len(nums)
	sum = nums[0] + nums[1] + nums[2]

	for i = 0; i  n-2; i++ {
		left, right = i+1, n-1
		for left  right {
			currsum = nums[i] + nums[left] + nums[right]
			if math.Abs(float64(currsum-target))  math.Abs(float64(sum-target)) {
				sum = currsum
			}
			if currsum  target {
				left++
			} else if currsum  target {
				right--
			} else {
				return currsum
			}
		}
	}

	return sum
}