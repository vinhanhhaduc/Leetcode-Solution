type Pair struct {
	node int
	prob float64
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	graph := make([][]Pair, n)
	for i, edge := range edges {
		u, v := edge[0], edge[1]
		prob := succProb[i]
		graph[u] = append(graph[u], Pair{v, prob})
		graph[v] = append(graph[v], Pair{u, prob})
	}
	a := make([]float64, n)
	for i := 0; i < n; i++ {
		a[i] = 0.0
	}
	a[start] = 1.0

	queue := []Pair{{start, 1.0}}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		currNode := current.node
		currProb := current.prob
		for _, neighbor := range graph[currNode] {
			nextNode := neighbor.node
			edgeProb := neighbor.prob
			newProb := currProb * edgeProb
			if newProb > a[nextNode] {
				a[nextNode] = newProb
				queue = append(queue, Pair{nextNode, newProb})
			}
		}
	}
	return a[end]
}