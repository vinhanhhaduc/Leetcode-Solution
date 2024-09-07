func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
func maxPoints(points [][]int) int {
	n := len(points)
	if n < 3 {
		return n
	}

	result := 1

	for i := 0; i < n; i++ {
		a := make(map[string]int)
		d := 0
		maxa := 0

		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			dx := points[j][0] - points[i][0]
			dy := points[j][1] - points[i][1]
			if dx == 0 && dy == 0 {
				d++
				continue
			}
			g := gcd(dx, dy)
			dx /= g
			dy /= g
			if dx < 0 {
				dx = -dx
				dy = -dy
			}
			aa := fmt.Sprintf("%d/%d", dy, dx)
			a[aa]++
			maxa = int(math.Max(float64(maxa), float64(a[aa])))
		}

	
		result = int(math.Max(float64(result), float64(maxa+d+1)))
	}

	return result
}