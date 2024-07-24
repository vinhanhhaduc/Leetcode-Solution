func helper(num int, mapping []int) int {
	str := strconv.Itoa(num)
	var mapNum strings.Builder

	for _, char := range str {
		digit := int(char - '0')
		mapNum.WriteString(strconv.Itoa(mapping[digit]))
	}

	mapVal, _ := strconv.Atoi(mapNum.String())
	return mapVal
}
func sortJumbled(mapping []int, nums []int) []int {
    mapVal := make(map[int]int, len(nums))

	for _, num := range nums {
		mapVal[num] = helper(num, mapping)
	}

	sort.SliceStable(nums, func(i, j int) bool {
		return mapVal[nums[i]] < mapVal[nums[j]]
	})

	return nums
}