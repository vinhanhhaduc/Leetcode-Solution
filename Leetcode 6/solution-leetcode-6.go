func convert(s string, numRows int) string {
    if numRows == 1 || numRows >= len(s) {
		return s
	}
	rows := make([]string, numRows)
	currR, direction := 0, -1
	for i := 0; i < len(s); i++ {
		rows[currR] += string(s[i])
		if currR == 0 || currR == numRows-1 {
			direction *= -1
		}
		currR += direction
	}

	var res strings.Builder
	for _, row := range rows {
		res.WriteString(row)
	}

	return res.String()
}