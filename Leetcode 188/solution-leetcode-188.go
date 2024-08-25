func maxProfit(k int, prices []int) int {
    if len(prices) == 0 || k == 0 {
        return 0
    }

    n := len(prices)
    if k >= n / 2 {
        maxp := 0
        for i := 1; i < n; i++ {
            if prices[i] > prices[i-1] {
                maxp += prices[i] - prices[i-1]
            }
        }
        return maxp
    }
    dp := make([][]int, k + 1)
    for i := 0; i <= k; i++ {
        dp[i] = make([]int, n)
    }

    for i := 1; i <= k; i++ {
        maxd := -prices[0]
        for j := 1; j < n; j++ {
            dp[i][j] = int(math.Max(float64(dp[i][j - 1]), float64(prices[j] + maxd)))
            maxd = int(math.Max(float64(maxd), float64(dp[i-1][j] - prices[j])))
        }
    }

    return dp[k][n - 1]
}