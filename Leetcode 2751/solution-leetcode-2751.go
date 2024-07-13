func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	n := len(positions)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = i
	}

	sort.Slice(a, func(i, j int) bool {
		return positions[a[i]] < positions[a[j]]
	})

	stack := []int{}

	for _, i := range a {
		if directions[i] == 'R' {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && healths[i] > 0 {
				j := stack[len(stack) - 1]
				stack = stack[:len(stack) - 1]

				if healths[j] > healths[i] {
					healths[j]--
					healths[i] = 0
					stack = append(stack, j)
				} else if healths[j] < healths[i] {
					healths[i]--
					healths[j] = 0
				} else {
					healths[i] = 0
					healths[j] = 0
				}
			}
		}
	}
	res := []int{}
	for i := 0; i < n; i++ {
		if healths[i] > 0 {
			res = append(res, healths[i])
		}
	}

	return res
}