func dfs(node int, graph map[int][]int, visited map[int]bool) {
	visited[node] = true
	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			dfs(neighbor, graph, visited)
		}
	}
}
func removeStones(stones [][]int) int {
	n := len(stones)
	visited := make(map[int]bool)
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if stones[i][0] == stones[j][0] || stones[i][1] == stones[j][1] {
				graph[i] = append(graph[i], j)
				graph[j] = append(graph[j], i)
			}
		}
	}

	m := 0
	for i := 0; i < n; i++ {
		if !visited[i] {
			dfs(i, graph, visited)
			m++
		}
	}
	return n - m
}