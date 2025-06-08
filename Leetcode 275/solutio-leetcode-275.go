func hIndex(citations []int) int {
	csize := len(citations)

	if citations[csize - 1] == 0 {
		return 0
	} else if csize == 1 {
		return 1
	}

	index := 0

	for i := 1; i <= csize; i++  {
		if citations[csize - i] >= i {
			index = i
		} else {
			break
		}
	}

	return index
}