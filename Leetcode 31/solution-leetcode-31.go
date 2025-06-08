package leetcode

func nextPermutation(nums []int) {
	i, j := 0, 0
	for i = len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		for j = len(nums) - 1; j > i; j-- {
			if nums[j] > nums[i] {
				break
			}
		}
		swap(&nums, i, j)
	}
	reverse(&nums, i+1, len(nums)-1)
}

func reverse(nums *[]int, i, j int) {
	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func swap(nums *[]int, i, j int) {
	(*nums)[i], (*nums)[j] = (*nums)[j], (*nums)[i]
}
func nextPermutation1(nums []int) {
	var n = len(nums)
	var pIdx = checkPermutationPossibility(nums)
	if pIdx == -1 {
		reverse(&nums, 0, n-1)
		return
	}

	var rp = len(nums) - 1
	for rp > 0 {
		if nums[rp] > nums[pIdx] {
			swap(&nums, pIdx, rp)
			break
		} else {
			rp--
		}
	}
	reverse(&nums, pIdx+1, n-1)
}

func checkPermutationPossibility(nums []int) (idx int) {
	var rp = len(nums) - 1
	for rp > 0 {
		if nums[rp-1] < nums[rp] {
			idx = rp - 1
			return idx
		}
		rp--
	}
	return -1
}