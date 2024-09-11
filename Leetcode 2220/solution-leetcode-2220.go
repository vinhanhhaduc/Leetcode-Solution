func minBitFlips(start int, goal int) int {
    a := start ^ goal
	res := 0
	for a > 0 {
		res += a & 1
		a >>= 1
	}
	
	return res
}