func maxCoins(nums []int) int {
    n := len(nums)
    pnums := make([]int, n+2)
    pnums[0], pnums[n+1] = 1, 1
    for i := 1; i <= n; i++ {
        pnums[i] = nums[i-1]
    }
    dp := make([][]int, n+2)
    for i := range dp {
        dp[i] = make([]int, n+2)
    }
    
    for length := 2; length <= n+1; length++ {
        for i := 0; i <= n+1-length; i++ {
            j := i + length
            for k := i + 1; k < j; k++ {
                dp[i][j] = max(dp[i][j], dp[i][k] + dp[k][j] + pnums[i]*pnums[k]*pnums[j])
            }
        }
    }
    
    return dp[0][n+1]
}