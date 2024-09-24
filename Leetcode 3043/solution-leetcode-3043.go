func longestCommonPrefix(arr1 []int, arr2 []int) int {
	a := make(map[string]struct{})
	for _, num := range arr1 {
		strNum := strconv.Itoa(num)
		for i := 1; i <= len(strNum); i++ {
			prefix := strNum[:i]
			a[prefix] = struct{}{}
		}
	}

	res := 0
	for _, num := range arr2 {
		strNum := strconv.Itoa(num)
		for i := 1; i <= len(strNum); i++ {
			prefix := strNum[:i]
			if _, found := a[prefix]; found {
				if i > res {
					res = i
				}
			} else {
				break
			}
		}
	}

	return res
}