func countSeniors(details []string) int {
    count := 0

	for _, detail := range details {
		age := detail[11:13]
		a, _ := strconv.Atoi(age)

		if a > 60 {
			count++
		}
	}

	return count
}