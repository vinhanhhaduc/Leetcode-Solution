func asteroidsDestroyed(mass int, asteroids []int) bool {
    sort.Ints(asteroids)
	total := mass

	for _, i := range asteroids {
		if i > total {
			return false
		} else {
			total += i
		}
	}

	return true
}