func helper(nums []int, mid int) int {
	count, left := 0, 0
	for right := 0; right < len(nums); right++ {
		for nums[right] - nums[left] > mid {
			left++
		}
		count += right - left
	}
	return count
}
func smallestDistancePair(nums []int, k int) int {
    sort.Ints(nums)
    l, r := 0, nums[len(nums) - 1]
	
	for l < r {
		mid := (l + r) / 2
		if helper(nums, mid) >= k {
			r = mid
		} else {
			l = mid + 1
		}
	}
	
	return l
}