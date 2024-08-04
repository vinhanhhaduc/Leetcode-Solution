package main

import (
	"fmt"
)

func check(x, y, N, M int) bool {
	return x >= 0 && y >= 0 && x < N && y < M
}
func dfs(board [][]rune, visited [][]bool, x, y int) int {
	directions := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	stack := [][2]int{{x, y}}
	count := 0

	for len(stack) > 0 {
		point := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if visited[point[0]][point[1]] {
			continue
		}
		visited[point[0]][point[1]] = true
		count++

		for _, dir := range directions {
			nx, ny := point[0] + dir[0], point[1] + dir[1]
			if check(nx, ny, len(board), len(board[0])) && board[nx][ny] == '#' && !visited[nx][ny] {
				stack = append(stack, [2]int{nx, ny})
			}
		}
	}
	return count
}

func solution(B []string) (int, int, int) {
	N := len(B)
	M := len(B[0])
	board := make([][]rune, N)
	visited := make([][]bool, N)

	for i := range B {
		board[i] = []rune(B[i])
		visited[i] = make([]bool, M)
	}

	res1, res2, res3 := 0, 0, 0

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if board[i][j] == '#' && !visited[i][j] {
				s := dfs(board, visited, i, j)
				switch s {
				case 1:
					res1++
				case 2:
					res2++
				case 3:
					res3++
				}
			}
		}
	}

	return res1, res2, res3
}

func main() {
	B := []string{
		"...",
		"....",
		"...",
	}
	res1, res2, res3 := solution(B)
	res := []int{res1, res2, res3}
	fmt.Println(res)
}
