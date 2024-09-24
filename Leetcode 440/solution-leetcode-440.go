func findKthNumber(n int, k int) int {
    helper := func(curr, n int) int {
		count := 0
		first, last := curr, curr
		for first <= n {
			count += min(n+1, last+1) - first
			first *= 10
			last = last*10 + 9
		}
		return count
	}
	curr := 1
	k--

	for k > 0 {
		count := helper(curr, n)
		if k >= count {
			k -= count
			curr++
		} else {
			k--
			curr *= 10
		}
	}

	return curr
}