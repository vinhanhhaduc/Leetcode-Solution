func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
    directions := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	
	res := make([][]int, 0)
	visited := make(map[[2]int]bool)
	
	r, c, dir := rStart, cStart, 0
	steps := 1
	res = append(res, []int{r, c})
	visited[[2]int{r, c}] = true

	for len(res) < rows * cols {
		for i := 0; i < 2; i++ {
			for j := 0; j < steps; j++ {
				r += directions[dir][0]
				c += directions[dir][1]
				if r >= 0 && r < rows && c >= 0 && c < cols {
					res = append(res, []int{r, c})
					visited[[2]int{r, c}] = true
				}
			}
			dir = (dir + 1) % 4
		}
		steps++
	}
	
	return res
}