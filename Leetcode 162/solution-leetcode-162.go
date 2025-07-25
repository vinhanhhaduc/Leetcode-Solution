func findPeakElement(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return 0
	}
	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if (mid == len(nums)-1 && nums[mid-1] < nums[mid]) || (mid > 0 && nums[mid-1] < nums[mid] && (mid <= len(nums)-2 && nums[mid+1] < nums[mid])) || (mid == 0 && nums[1] < nums[0]) {
			return mid
		}
		if mid > 0 && nums[mid-1] < nums[mid] {
			low = mid + 1
		}
		if mid > 0 && nums[mid-1] > nums[mid] {
			high = mid - 1
		}
		if mid == low {
			low++
		}
		if mid == high {
			high--
		}
	}
	return -1
}
func findPeakElement1(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		mid := low + (high-low)>>1
		if nums[mid] > nums[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}