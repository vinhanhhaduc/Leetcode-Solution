func minimumTotal(triangle [][]int) int {
    if len(triangle) == 0 {
		return 0
	}
	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col < len(triangle[row]); col++ {
			triangle[row][col] += int(math.Min(float64(triangle[row+1][col]), float64(triangle[row+1][col+1])))
		}
	}
	return triangle[0][0]
}