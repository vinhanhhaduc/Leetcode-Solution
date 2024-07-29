func hasAllCodes(s string, k int) bool {
    t := 1 << k
	a := make(map[string]struct{})
	for i := 0; i <= len(s)-k; i++ {
		st := s[i : i+k]
		a[st] = struct{}{}
	}
	return len(a) == t
}