package main

import (
	"fmt"
	"sort"
)

func solution(P []int, S []int, N int) int {
	sort.Slice(S, func(i, j int) bool {
		return S[i] > S[j]
	})
	total := 0
	for _, i := range P {
		total += i
	}
	res := 0
	k := 0
	for i := 0; i < N; i++ {
		k += S[i]
		res++
		if k >= total {
			break
		}
	}

	return res
}

func main() {
	P := []int{4,4,2,4}
	S := []int{5,5,2,5}
	N := len(P)
	fmt.Println(solution(P, S, N))
}
