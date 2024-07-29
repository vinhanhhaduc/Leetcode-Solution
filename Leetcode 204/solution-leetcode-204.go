func countPrimes(n int) int {
    if n <= 2 {
        return 0
    }
    san := make([]bool, n)
    count := 0
    for i := 2; i < n; i++ {
		san[i] = true
	}
    for i := 2; i * i < n; i++ {
        if san[i] {
            for j := i * i; j < n; j += i {
                san[j] = false
            }
        }
    }
    for i := 2; i < n; i++ {
		if san[i] {
			count++
		}
	}

	return count
}