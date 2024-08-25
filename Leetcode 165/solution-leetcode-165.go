func compareVersion(version1 string, version2 string) int {
	revisions1 := strings.Split(version1, ".")
	revisions2 := strings.Split(version2, ".")
	maxLength := max(len(revisions1), len(revisions2))
	for i := 0; i < maxLength; i++ {
		rev1 := 0
		rev2 := 0
		if i < len(revisions1) {
			rev1, _ = strconv.Atoi(revisions1[i])
		}
		if i < len(revisions2) {
			rev2, _ = strconv.Atoi(revisions2[i])
		}
		if rev1 > rev2 {
			return 1
		} else if rev1 < rev2 {
			return -1
		}
	}

	return 0
}
