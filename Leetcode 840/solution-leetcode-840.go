func numMagicSquaresInside(grid [][]int) int {
    count := 0
	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[0])-2; j++ {
			if check(grid, i, j) {
				count++
			}
		}
	}
	return count
}
func check(grid [][]int, row, col int) bool {
	a := make([]bool, 10)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := grid[row+i][col+j]
			if num < 1 || num > 9 || a[num] {
				return false
			}
			a[num] = true
		}
	}
	sum := grid[row][col] + grid[row][col + 1] + grid[row][col + 2]
	return grid[row + 1][col] + grid[row + 1][col + 1]+grid[row + 1][col + 2] == sum &&
		grid[row + 2][col] + grid[row + 2][col + 1] + grid[row + 2][col + 2] == sum &&
		grid[row][col] + grid[row + 1][col] + grid[row + 2][col] == sum &&
		grid[row][col + 1] + grid[row + 1][col+1] + grid[row + 2][col + 1] == sum &&
		grid[row][col + 2] + grid[row + 1][col + 2] + grid[row + 2][col + 2] == sum &&
		grid[row][col] + grid[row + 1][col + 1] + grid[row + 2][col + 2] == sum &&
		grid[row][col + 2] + grid[row + 1][col + 1] + grid[row + 2][col] == sum
}