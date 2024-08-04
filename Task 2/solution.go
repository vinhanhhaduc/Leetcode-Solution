package main

import (
	"fmt"
)
func solution(D []int, N int, X int) int {
	days := 1
	mind := D[0]
	maxd := D[0]
	for i := 1; i < N; i++ {
		if D[i] < mind {
			mind = D[i]
		}
		if D[i] > maxd {
			maxd = D[i]
		}
		if maxd - mind > X {
			days++
			mind = D[i]
			maxd = D[i]
		}
	}
	
	return days
}

func main() {
	D := []int{1, 12, 10, 4, 5, 2}
	N := len(D1)
	X := 2
	fmt.Println(solution(D, N, X))
}