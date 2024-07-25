func solve(l []int, r []int) []int {
	res := make([]int, 0, len(l) + len(r))
	i, j := 0, 0

	for i < len(l) && j < len(r) {
		if l[i] <= r[j] {
			res = append(res, l[i])
			i++
		} else {
			res = append(res, r[j])
			j++
		}
	}

	res = append(res, l[i:]...)
	res = append(res, r[j:]...)

	return res
}
func sortArray(nums []int) []int {
    if len(nums) < 2 {
		return nums
	}

	mid := len(nums) / 2
	l := sortArray(nums[:mid])
	r := sortArray(nums[mid:])

	return solve(l, r)
}