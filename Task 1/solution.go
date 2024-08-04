package main

import (
	"fmt"
	"strconv"
)
func helper(num int) (int, int) {
	n := strconv.Itoa(num)
	firstDigit, _ := strconv.Atoi(string(n[0]))
	lastDigit, _ := strconv.Atoi(string(n[len(n)-1]))
	return firstDigit, lastDigit
}

func solution(numbers []int, N int) int {
	ldmap := make(map[int]int)
	fdmap := make(map[int]int)

	for _, num := range numbers {
		firstDigit, lastDigit := helper(num)
		ldmap[lastDigit]++
		fdmap[firstDigit]++
	}
	count := 0
	for _, num := range numbers {
		firstDigit, _ := helper(num)
		if ldmap[firstDigit] > 0 {
			count += ldmap[firstDigit]
		}
	}

	return count
}
func main() {
	numbers := []int{30, 12, 29, 91}
	fmt.Println(solution(numbers, len(numbers)))

}