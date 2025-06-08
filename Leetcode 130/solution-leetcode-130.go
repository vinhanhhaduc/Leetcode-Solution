var dir = [][]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	m, n := len(board[0]), len(board)
	uf := template.UnionFind{}
	uf.Init(n*m + 1)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if (i == 0 || i == n-1 || j == 0 || j == m-1) && board[i][j] == 'O' {
				uf.Union(i*m+j, n*m)
			} else if board[i][j] == 'O' {
				if board[i-1][j] == 'O' {
					uf.Union(i*m+j, (i-1)*m+j)
				}
				if board[i+1][j] == 'O' {
					uf.Union(i*m+j, (i+1)*m+j)
				}
				if board[i][j-1] == 'O' {
					uf.Union(i*m+j, i*m+j-1)
				}
				if board[i][j+1] == 'O' {
					uf.Union(i*m+j, i*m+j+1)
				}

			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if uf.Find(i*m+j) != uf.Find(n*m) {
				board[i][j] = 'X'
			}
		}
	}
}
func solve1(board [][]byte) {
	for i := range board {
		for j := range board[i] {
			if i == 0 || i == len(board)-1 || j == 0 || j == len(board[i])-1 {
				if board[i][j] == 'O' {
					dfs130(i, j, board)
				}
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == '*' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func dfs130(i, j int, board [][]byte) {
	if i < 0 || i > len(board)-1 || j < 0 || j > len(board[i])-1 {
		return
	}
	if board[i][j] == 'O' {
		board[i][j] = '*'
		for k := 0; k < 4; k++ {
			dfs130(i+dir[k][0], j+dir[k][1], board)
		}
	}
}