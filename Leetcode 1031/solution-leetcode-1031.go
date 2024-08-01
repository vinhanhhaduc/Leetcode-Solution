func maxSumTwoNoOverlap(nums []int, firstLen int, secondLen int) int {
	return max(helper(nums, firstLen, secondLen), helper(nums, secondLen, firstLen))
}

func helper(nums []int, l int, m int) int {
	n := len(nums)
	pref := make([]int, n+1)
	for i := 0; i < n; i++ {
		pref[i+1] = pref[i] + nums[i]
	}
	
	maxl, maxm := 0, 0
	res := 0
	for i := l + m; i <= n; i++ {
		maxl = max(maxl, pref[i-m] - pref[i-m-l])
		res = max(res, maxl + pref[i] - pref[i-m])
	}
	
	for i := l + m; i <= n; i++ {
		maxm = max(maxm, pref[i-l] - pref[i-l-m])
		res = max(res, maxm + pref[i] - pref[i-l])
	}
	
	return res
}