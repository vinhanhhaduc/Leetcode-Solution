func luckyNumbers (matrix [][]int) []int {
    m := len(matrix)
	n := len(matrix[0])
	res := []int{}
	
	rmin := make([]int, m)
	for i := 0; i < m; i++ {
		rmin[i] = matrix[i][0]
		for j := 1; j < n; j++ {
			if matrix[i][j] < rmin[i] {
				rmin[i] = matrix[i][j]
			}
		}
	}
	cmax := make([]int, n)
	for j := 0; j < n; j++ {
		cmax[j] = matrix[0][j]
		for i := 1; i < m; i++ {
			if matrix[i][j] > cmax[j] {
				cmax[j] = matrix[i][j]
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == rmin[i] && matrix[i][j] == cmax[j] {
				res = append(res, matrix[i][j])
			}
		}
	}

	return res
}