func addHeight(heightMap map[int]int, heights *[]int, h int) {
	heightMap[h]++
	if heightMap[h] == 1 { // If height is added for the first time
		*heights = append(*heights, h)
		sort.Ints(*heights)
	}
}
func removeHeight(heightMap map[int]int, heights *[]int, h int) {
	heightMap[h]--
	if heightMap[h] == 0 {
		delete(heightMap, h)
		for i := range *heights {
			if (*heights)[i] == h {
				*heights = append((*heights)[:i], (*heights)[i+1:]...)
				break
			}
		}
	}
}
func getSkyline(buildings [][]int) [][]int {
	var res [][]int
	var points []struct {
		x, height int
	}
	heightMap := map[int]int{0: 1}
	var heights []int
	for _, build := range buildings {
		points = append(points, struct{ x, height int }{build[0], -build[2]}) 
		points = append(points, struct{ x, height int }{build[1], build[2]}) 
	}
	sort.Slice(points, func(i, j int) bool {
		if points[i].x == points[j].x {
			return points[i].height < points[j].height
		}
		return points[i].x < points[j].x
	})

	ongoingHeight := 0
	for _, point := range points {
		currPoint := point.x
		heightAtCurrentPoint := point.height

		if heightAtCurrentPoint < 0 {
			addHeight(heightMap, &heights, -heightAtCurrentPoint)
		} else {
			removeHeight(heightMap, &heights, heightAtCurrentPoint)
		}
		var currMaxHeight int
		if len(heights) > 0 {
			currMaxHeight = heights[len(heights)-1]
		} else {
			currMaxHeight = 0
		}

		if ongoingHeight != currMaxHeight {
			ongoingHeight = currMaxHeight
			res = append(res, []int{currPoint, ongoingHeight})
		}
	}

	return res
}