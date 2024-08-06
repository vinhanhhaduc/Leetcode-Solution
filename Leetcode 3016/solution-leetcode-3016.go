func minimumPushes (word string) int {
	freq := make([]int, 26)
	for _, c := range word {
		freq[int(c - 'a')]++
	}
	sort.Sort(sort.Reverse(sort.IntSlice(freq)))

	res := 0
	mult := 1
	count := 8
	for i := 0; i < 26; i++ {
		if freq[i] == 0 {
			break
		}
		res += freq[i] * mult
		count--
		if count == 0 {
			mult++
			count = 8
		}
	}

	return res
}