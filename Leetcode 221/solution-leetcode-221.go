func maximalSquare(matrix [][]byte) int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	
	res := 0
	for i := 0; i < m; i++ {
		dp[i][0] = int(matrix[i][0] - '0')
		if dp[i][0] > res {
			res = dp[i][0]
		}
	}
	for j := 0; j < n; j++ {
		dp[0][j] = int(matrix[0][j] - '0')
		if dp[0][j] > res {
			res = dp[0][j]
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
				if dp[i][j] > res {
					res = dp[i][j]
				}
			}
		}
	}
	
	return res * res
}