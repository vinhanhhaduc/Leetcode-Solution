func timeToMinutes(time string) int {
	parts := strings.Split(time, ":")
	hours, _ := strconv.Atoi(parts[0])
	minutes, _ := strconv.Atoi(parts[1])
	return hours * 60 + minutes
}
func findMinDifference(timePoints []string) int {
    minDiff := 1440 
    if len(timePoints) > 1440 { 
		return 0
	}
	minutes := make([]int, len(timePoints))
	for i, time := range timePoints {
		minutes[i] = timeToMinutes(time)
	}

	sort.Ints(minutes)

	for i := 1; i < len(minutes); i++ {
		minDiff = min(minDiff, minutes[i]-minutes[i-1])
	}

	circularDiff := 1440 + minutes[0] - minutes[len(minutes)-1]
	minDiff = min(minDiff, circularDiff)

	return minDiff
}