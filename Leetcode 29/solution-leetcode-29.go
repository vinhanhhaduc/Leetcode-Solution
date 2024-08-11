func divide(dividend int, divisor int) int {
    if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	sign := 1
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		sign = -1
	}
	dividendAbs := int(math.Abs(float64(dividend)))
	divisorAbs := int(math.Abs(float64(divisor)))

	res := 0
	for dividendAbs >= divisorAbs {
		temp := divisorAbs
		multiple := 1
		for dividendAbs >= (temp << 1) {
			temp <<= 1
			multiple <<= 1
		}
		dividendAbs -= temp
		res += multiple
	}
	res *= sign
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	if res < math.MinInt32 {
		return math.MinInt32
	}

	return res
}