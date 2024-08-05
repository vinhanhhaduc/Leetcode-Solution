package main

import (
	"fmt"
)

func solution(A []int) int {
	n := len(A)
	if n < 2 {
		return 0
	}
	res := 0
	s := make([]bool, n)

	for i := 0; i < n; i++ {
		curr := A[i]
		next := A[(i + 1) % n]
		if !s[i] && !s[(i + 1) % n] && (curr + next) % 2 == 0 {
			res++
			s[i] = true
			s[(i + 1) % n] = true
		}
	}

	return res
}

func main() {
	A := []int{5,5,5,5,5,5}
	fmt.Println(solution(A))
}
