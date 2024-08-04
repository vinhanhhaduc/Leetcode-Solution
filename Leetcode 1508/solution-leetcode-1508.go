const MOD = 1_000_000_007
func rangeSum(nums []int, n int, left int, right int) int {
	pref := make([]int, n + 1)
	for i := 0; i < n; i++ {
		pref[i + 1] = pref[i] + nums[i]
	}
	var sums []int
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			sum := pref[j] - pref[i]
			sums = append(sums, sum)
		}
	}

	sort.Ints(sums)
	res := 0
	for i := left - 1; i < right; i++ {
		res = (res + sums[i]) % MOD
	}

	return res
}
