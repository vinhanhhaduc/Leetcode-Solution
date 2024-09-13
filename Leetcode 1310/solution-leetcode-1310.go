func xorQueries(arr []int, queries [][]int) []int {
    n := len(arr)
	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] ^ arr[i-1]
	}
	result := make([]int, len(queries))
	for i, query := range queries {
		left, right := query[0], query[1]
		result[i] = prefix[right+1] ^ prefix[left]
	}

	return result
}