const INF = int(1e9 + 5)

type Edge struct {
	to, weight int
}

type MinHeap [][]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func dijkstra(adj [][]Edge, src, dest, n int) int {
	dist := make([]int, n)
	for i := range dist {
		dist[i] = INF
	}
	dist[src] = 0

	minHeap := &MinHeap{}
	heap.Init(minHeap)
	heap.Push(minHeap, []int{0, src})

	for minHeap.Len() > 0 {
		current := heap.Pop(minHeap).([]int)
		d, u := current[0], current[1]

		if dist[u] != d {
			continue
		}
		if u == dest {
			return d
		}

		for _, neighbor := range adj[u] {
			v, w := neighbor.to, neighbor.weight
			if d+w < dist[v] {
				dist[v] = d + w
				heap.Push(minHeap, []int{dist[v], v})
			}
		}
	}

	return INF
}

func modifiedGraphEdges(n int, edges [][]int, source int, destination int, target int) [][]int {
	adj := make([][]Edge, n)
	negEdges := [][]int{}

	for _, edge := range edges {
		u, v, w := edge[0], edge[1], edge[2]
		if w == -1 {
			negEdges = append(negEdges, []int{u, v, -1})
			continue
		}
		adj[u] = append(adj[u], Edge{v, w})
		adj[v] = append(adj[v], Edge{u, w})
	}
	initialDist := dijkstra(adj, source, destination, n)
	if initialDist < target {
		return [][]int{}
	}
	ok := initialDist == target
	for _, edge := range negEdges {
		u, v := edge[0], edge[1]
		if ok {
			edge[2] = INF
		} else {
			edge[2] = 1
			adj[u] = append(adj[u], Edge{v, 1})
			adj[v] = append(adj[v], Edge{u, 1})

			newDist := dijkstra(adj, source, destination, n)
			if newDist <= target {
				ok = true
				edge[2] += target - newDist
			}
		}
	}

	if !ok {
		return [][]int{}
	}
	for i := range edges {
		if edges[i][2] == -1 {
			for _, negEdge := range negEdges {
				if edges[i][0] == negEdge[0] && edges[i][1] == negEdge[1] {
					edges[i][2] = negEdge[2]
					break
				}
			}
		}
	}

	return edges
}