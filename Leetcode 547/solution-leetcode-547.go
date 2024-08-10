func dfs(city int, n int, isConnected [][]int, visited []bool) {
	for i := 0; i < n; i++ {
		if isConnected[city][i] == 1 && !visited[i] {
			visited[i] = true
			dfs(i, n, isConnected, visited)
		}
	}
}
func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
	visited := make([]bool, n)
	count := 0

	for i := 0; i < n; i++ {
		if !visited[i] {
			visited[i] = true
			dfs(i, n, isConnected, visited)
			count++
		}
	}

	return count
}