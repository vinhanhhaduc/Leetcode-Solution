func floydWarshall(dist [][]int) {
	for k := 0; k < 26; k++ {
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				if dist[i][k] < math.MaxInt64 && dist[k][j] < math.MaxInt64 {
					if dist[i][j] > dist[i][k]+dist[k][j] {
						dist[i][j] = dist[i][k] + dist[k][j]
					}
				}
			}
		}
	}
}
func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	n := len(source)
	if n != len(target) {
		return -1
	}
	dist := make([][]int, 26)
	for i := range dist {
		dist[i] = make([]int, 26)
		for j := range dist[i] {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = math.MaxInt64
			}
		}
	}

	for i := 0; i < len(original); i++ {
		o, c := original[i]-'a', changed[i]-'a'
		if dist[o][c] > cost[i] {
			dist[o][c] = cost[i]
		}
	}

	floydWarshall(dist)

	totalCost := 0
	for i := 0; i < n; i++ {
		if source[i] != target[i] {
			s, t := source[i]-'a', target[i]-'a'
			if dist[s][t] == math.MaxInt64 {
				return -1
			}
			totalCost += dist[s][t]
		}
	}

	return int64(totalCost)
}