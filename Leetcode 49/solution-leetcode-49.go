func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}
func groupAnagrams(strs []string) [][]string {
    a := make(map[string][]string)
    for _, val := range strs {
        sortStr := sortString(val)
        a[sortStr] = append(a[sortStr], val)
    }
    res := make([][]string, 0, len(a))
	for _, val := range a {
		res = append(res, val)
	}

	return res
}