func isRectangleCover(rectangles [][]int) bool {
    minX, minY := 1<<31-1, 1<<31-1
	maxX, maxY := -1<<31, -1<<31
	area := 0
	corners := make(map[[2]int]int)
	for _, rect := range rectangles {
		x1, y1, x2, y2 := rect[0], rect[1], rect[2], rect[3]
		if x1 < minX {
			minX = x1
		}
		if y1 < minY {
			minY = y1
		}
		if x2 > maxX {
			maxX = x2
		}
		if y2 > maxY {
			maxY = y2
		}
		area += (x2 - x1) * (y2 - y1)

		// Step 3: Update corner counts
		corners[[2]int{x1, y1}]++
		corners[[2]int{x1, y2}]++
		corners[[2]int{x2, y1}]++
		corners[[2]int{x2, y2}]++
	}
	ea := (maxX - minX) * (maxY - minY)
	if area != ea {
		return false
	}
	ec := [][2]int{{minX, minY}, {minX, maxY}, {maxX, minY}, {maxX, maxY}}

	for _, corner := range ec {
		if corners[corner] != 1 {
			return false
		}
		delete(corners, corner)
	}
	for _, count := range corners {
		if count != 2 && count != 4 {
			return false
		}
	}

	return true
}