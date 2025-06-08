func majorityElement229(nums []int) []int {
	count1, count2, candidate1, candidate2 := 0, 0, 0, 1
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		} else if count1 <= 0 {
			candidate1, count1 = num, 1
		} else if count2 <= 0 {
			candidate2, count2 = num, 1
		} else {
			count1--
			count2--
		}
	}
	count1, count2 = 0, 0
	for _, num := range nums {
		if num == candidate1 {
			count1++
		} else if num == candidate2 {
			count2++
		}
	}
	length := len(nums)
	if count1 > length/3 && count2 > length/3 {
		return []int{candidate1, candidate2}
	}
	if count1 > length/3 {
		return []int{candidate1}
	}
	if count2 > length/3 {
		return []int{candidate2}
	}
	return []int{}
}

func majorityElement229_1(nums []int) []int {
	result, m := make([]int, 0), make(map[int]int)
	for _, val := range nums {
		if v, ok := m[val]; ok {
			m[val] = v + 1
		} else {
			m[val] = 1
		}
	}
	for k, v := range m {
		if v > len(nums)/3 {
			result = append(result, k)
		}
	}
	return result
}