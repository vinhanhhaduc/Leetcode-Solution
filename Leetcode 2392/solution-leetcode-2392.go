func topologicalSort(k int, conditions [][]int) ([]int, bool) {
	inDegree := make([]int, k+1)
	graph := make(map[int][]int)

	for _, cond := range conditions {
		u, v := cond[0], cond[1]
		graph[u] = append(graph[u], v)
		inDegree[v]++
	}

	queue := []int{}
	for i := 1; i <= k; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	sortedOrder := []int{}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		sortedOrder = append(sortedOrder, node)

		for _, neighbor := range graph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if len(sortedOrder) != k {
		return nil, false
	}

	return sortedOrder, true
}
func buildMatrix(k int, rowc [][]int, colc [][]int) [][]int {
    ro, rok := topologicalSort(k, rowc)
	co, cok := topologicalSort(k, colc)

	if !rok || !cok {
		return [][]int{}
	}

	pos := make(map[int][2]int)
	for i, num := range ro {
		pos[num] = [2]int{i, pos[num][1]}
	}
	for j, num := range co {
		pos[num] = [2]int{pos[num][0], j}
	}

	matrix := make([][]int, k)
	for i := range matrix {
		matrix[i] = make([]int, k)
	}

	for i, j := range pos {
		matrix[j[0]][j[1]] = i
	}

	return matrix
}