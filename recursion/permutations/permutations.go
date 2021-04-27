package permutations

func permute(nums []int) [][]int {
	return backtrack(nums, 0)
}

func backtrack(nums []int, first int) [][]int {
	if first == len(nums) {
		cpy := make([]int, len(nums))
		copy(cpy, nums)
		return [][]int{cpy}
	}

	var res [][]int
	for i := first; i < len(nums); i++ {
		nums[first], nums[i] = nums[i], nums[first]
		res = append(res, backtrack(nums, first+1)...)
		nums[first], nums[i] = nums[i], nums[first]
	}

	return res
}
