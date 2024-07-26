func findTheCity(n int, edges [][]int, distanceThreshold int) int {
    dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			if i != j {
				dist[i][j] = math.MaxInt
			}
		}
	}
	for _, edge := range edges {
		from, to, weight := edge[0], edge[1], edge[2]
		dist[from][to] = weight
		dist[to][from] = weight
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dist[i][k] != math.MaxInt && dist[k][j] != math.MaxInt {
					dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][j])
				}
			}
		}
	}
	minr := math.MaxInt
	res := -1

	for i := 0; i < n; i++ {
		countr := 0
		for j := 0; j < n; j++ {
			if i != j && dist[i][j] <= distanceThreshold {
				countr++
			}
		}

		if countr <= minr {
			if countr < minr || i > res {
				minr = countr
				res = i
			}
		}
	}

	return res
}