func fallingSquares(positions [][]int) []int {
    n := len(positions)
	heights := make([]int, n)
	res := make([]int, n)  
	maxH := 0 

	for i := 0; i < n; i++ {
		lefti, sideLengthi := positions[i][0], positions[i][1]
		righti := lefti + sideLengthi
		heights[i] = sideLengthi
		for j := 0; j < i; j++ {
			leftj, sideLengthj := positions[j][0], positions[j][1]
			rightj := leftj + sideLengthj
			if lefti < rightj && righti > leftj {
				heights[i] = int(math.Max(float64(heights[i]), float64(heights[j]+sideLengthi)))
			}
		}
		maxH = int(math.Max(float64(maxH), float64(heights[i])))
		res[i] = maxH
	}

	return res
}