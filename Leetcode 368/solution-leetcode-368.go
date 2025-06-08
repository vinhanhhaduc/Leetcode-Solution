func lexicalOrder(n int) []int {
	var ans []int

	var solve func(i int)
	solve = func(i int) {
		if i > n {
			return
		}

		cur := i
		ans = append(ans, cur)
		for j := 0; j <= 9; j++ {
			tmp := strconv.Itoa(cur) + strconv.Itoa(j) // 1 0
			num, _ := strconv.Atoi(tmp)
			solve(num)
		}
	}

	for i := 1; i <= 9; i++ {
		solve(i)
	}

	return ans
}