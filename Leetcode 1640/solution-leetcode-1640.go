func canFormArray(arr []int, pieces [][]int) bool {
    a := make(map[int][]int)
	for _, piece := range pieces {
		a[piece[0]] = piece
	}

	i := 0
	for i < len(arr) {
		if piece, k := a[arr[i]]; k {
			for j := 0; j < len(piece); j++ {
				if arr[i + j] != piece[j] {
					return false
				}
			}
			i += len(piece)
		} else {
			return false
		}
	}
	return true
}
