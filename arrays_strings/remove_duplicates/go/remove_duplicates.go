package remove_duplicates

func RemoveDuplicates(nums []int) int {
	l := len(nums)
	for j := 0; j < l; j++ {
		for k := j+1; k < l; k++ {
			if nums[k] == nums[j] {
				_ = append(nums[:k], nums[k+1:]...)
				l--
				k--
			}
		}
	}

	return l
}
