func nthUglyNumber(n int) int {
    res := make([]int, n)
	res[0] = 1
	i2, i3, i5 := 0, 0, 0

	for i := 1; i < n; i++ {
		minU := min(res[i2] * 2, min(res[i3] * 3, res[i5] * 5))
		res[i] = minU

		if minU == res[i2] * 2 {
			i2++
		}
		if minU == res[i3] * 3 {
			i3++
		}
		if minU == res[i5] * 5 {
			i5++
		}
	}

	return res[n - 1]
}