func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	m, n := len(nums1), len(nums2)
	imin, imax, half := 0, m, (m + n + 1)/2

	for imin <= imax {
		i := (imin + imax) / 2
		j := half - i
		if i < m && nums1[i] < nums2[j - 1] {
			imin = i + 1
		} else if i > 0 && nums1[i - 1] > nums2[j] {
			imax = i - 1
		} else {
			maxOfLeft := 0
			if i == 0 {
				maxOfLeft = nums2[j - 1]
			} else if j == 0 {
				maxOfLeft = nums1[i - 1]
			} else {
				maxOfLeft = int(math.Max(float64(nums1[i - 1]), float64(nums2[j - 1])))
			}

			if (m+n)%2 == 1 {
				return float64(maxOfLeft)
			}

			minOfRight := 0
			if i == m {
				minOfRight = nums2[j]
			} else if j == n {
				minOfRight = nums1[i]
			} else {
				minOfRight = int(math.Min(float64(nums1[i]), float64(nums2[j])))
			}

			return float64(maxOfLeft + minOfRight) / 2.0
		}
	}

	return 0.0
}