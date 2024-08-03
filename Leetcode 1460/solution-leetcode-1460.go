func canBeEqual(target []int, arr []int) bool {
    if len(target) != len(arr) {
		return false
    }
	counts := make(map[int]int)
	for _, num := range target {
		counts[num]++
	}
	for _, num := range arr {
		counts[num]--
		if counts[num] == 0 {
			delete(counts, num)
		}
	}
	return len(counts) == 0
}