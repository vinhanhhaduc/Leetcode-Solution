type charCount struct {
	count int
	char  byte
}
func longestDiverseString(a int, b int, c int) string {
	result := []byte{}
	counts := []charCount{
		{a, 'a'},
		{b, 'b'},
		{c, 'c'},
	}

	for {
		sort.Slice(counts, func(i, j int) bool {
			return counts[i].count > counts[j].count
		})
		added := false
		for i := 0; i < 3; i++ {
			if counts[i].count == 0 {
				continue
			}
			n := len(result)
			if n >= 2 && result[n-1] == counts[i].char && result[n-2] == counts[i].char {
				continue
			}
			result = append(result, counts[i].char)
			counts[i].count--
			added = true
			break
		}
		if !added {
			break
		}
	}

	return string(result)
}