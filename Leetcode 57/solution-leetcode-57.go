func insert(intervals [][]int, newInterval []int) [][]int {
    var result [][]int
	i := 0
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = int(math.Min(float64(newInterval[0]), float64(intervals[i][0])))
		newInterval[1] = int(math.Max(float64(newInterval[1]), float64(intervals[i][1])))
		i++
	}
	result = append(result, newInterval)
	for i < len(intervals) {
		result = append(result, intervals[i])
		i++
	}

	return result
}