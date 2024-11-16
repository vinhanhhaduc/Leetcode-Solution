func countFairPairs(nums []int, lower int, upper int) int64 {
    sort.Ints(nums)
	count := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		l := sort.Search(n, func(j int) bool {
			return j > i && nums[i]+nums[j] >= lower
		})
		r := sort.Search(n, func(j int) bool {
			return j > i && nums[i]+nums[j] > upper
		})
		if l < r {
			count += r - l
		}
	}

	return int64(count)
}