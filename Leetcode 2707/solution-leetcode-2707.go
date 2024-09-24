func minExtraChar(s string, dictionary []string) int {
    a := make(map[string]bool)
	for _, word := range dictionary {
		a[word] = true
	}

	n := len(s)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[n] = 0 
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j <= n; j++ {
			substring := s[i:j]
			if a[substring] {
				dp[i] = min(dp[i], dp[j])
			}
		}
		dp[i] = min(dp[i], dp[i+1]+1)
	}

	return dp[0]
}