func helper(nums []int, t int) []int {
	stack := make([]int, 0, t)
	td := len(nums) - t
	for _, num := range nums {
		for len(stack) > 0 && stack[len(stack)-1] < num && td > 0 {
			stack = stack[:len(stack)-1]
			td--
		}
		stack = append(stack, num)
	}
	return stack[:t]
}

func merge(nums1, nums2 []int) []int {
	merged := make([]int, 0, len(nums1)+len(nums2))
	for len(nums1) > 0 || len(nums2) > 0 {
		if compare(nums1, nums2) {
			merged = append(merged, nums1[0])
			nums1 = nums1[1:]
		} else {
			merged = append(merged, nums2[0])
			nums2 = nums2[1:]
		}
	}
	return merged
}
func compare(nums1, nums2 []int) bool {
	for i := 0; i < len(nums1) && i < len(nums2); i++ {
		if nums1[i] > nums2[i] {
			return true
		}
		if nums1[i] < nums2[i] {
			return false
		}
	}
	return len(nums1) > len(nums2)
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	var res []int
	for i := 0; i <= k && i <= len(nums1); i++ {
		if k-i > len(nums2) {
			continue
		}
		s1 := helper(nums1, i)
		s2 := helper(nums2, k-i)
		merged := merge(s1, s2)
		if compare(merged, res) {
			res = merged
		}
	}
	return res
}
