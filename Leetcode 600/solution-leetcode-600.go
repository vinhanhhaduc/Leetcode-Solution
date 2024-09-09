func findIntegers(n int) int {
    dp := make([]int, 32)
    dp[0], dp[1] = 1, 2
    
    for i := 2; i < 32; i++ {
        dp[i] = dp[i-1] + dp[i-2]
    }
    pb := 0
    res := 0 
    for i := 30; i >= 0; i-- {
        if (n & (1 << i)) != 0 {
            res += dp[i]
            if pb == 1 {
                return res
            }
            pb = 1
        } else {
            pb = 0
        }
    }
    return res + 1
}