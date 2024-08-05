func kthDistinct(arr []string, k int) string {
    count := make(map[string]int)
	for _, i := range arr {
		count[i]++
	}
	var res []string
	for _, i := range arr {
		if count[i] == 1 {
			res = append(res, i)
		}
	}
	if k <= len(res) {
		return res[k - 1]
	}
	return ""
}