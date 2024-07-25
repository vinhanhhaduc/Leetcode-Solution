func check(board [][]byte, row, col int, char byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == char {
			return false
		}
		if board[i][col] == char {
			return false
		}
		if board[(row / 3) * 3 + i / 3][(col / 3) * 3 + i % 3] == char {
			return false
		}
	}
	return true
}
func solve(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for k := byte('1'); k <= '9'; k++ {
					if check(board, i, j, k) {
						board[i][j] = k
						if solve(board) {
							return true
						} else {
							board[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}
func solveSudoku(board [][]byte)  {
    solve(board)
}