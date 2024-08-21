func stoneGameII(piles []int) int {
    n := len(piles)
    suffixSum := make([]int, n)
    suffixSum[n-1] = piles[n-1]
    for i := n - 2; i >= 0; i-- {
        suffixSum[i] = piles[i] + suffixSum[i+1]
    }

    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n+1)
    }
    for i := n - 1; i >= 0; i-- {
        for m := 1; m <= n; m++ {
            if i + 2*m >= n {
                dp[i][m] = suffixSum[i]
            } else {
                for x := 1; x <= 2*m; x++ {
                    dp[i][m] = max(dp[i][m], suffixSum[i] - dp[i+x][max(m, x)])
                }
            }
        }
    }

    return dp[0][1]
}