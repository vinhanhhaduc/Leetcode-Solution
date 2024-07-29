func numTeams(rating []int) int {
    n := len(rating)
    count := 0
	for i := 1; i < n - 1; i++ {
		sl := 0
		ll := 0
		for j := 0; j < i; j++ {
			if rating[j] < rating[i] {
				sl++
			} else if rating[j] > rating[i] {
				ll++
			}
		}
		sr := 0
		lr := 0
		for k := i + 1; k < n; k++ {
			if rating[k] < rating[i] {
				sr++
			} else if rating[k] > rating[i] {
				lr++
			}
		}
		count += sl * lr
		count += ll * sr
	}

	return count
}