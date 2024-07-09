func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func averageWaitingTime(customers [][]int) float64 {
    n := 0
    m := 0
    length := len(customers)
    for _, customer := range customers {
        a := customer[0]
        b := customer[1]
        if max(m, a) == a {
            n += b
            m = a + b
        } else {
            m += b
            n += m - a 
        }
    }
    return float64(n) / float64(length)
}