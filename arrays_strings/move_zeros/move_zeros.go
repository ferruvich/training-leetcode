package move_zeros

func MoveZeroes(nums []int) {
	var zeros = 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			tmp := append(nums[i+1:], 0)
			_ = append(nums[:i], tmp...)
			zeros++

			if zeros+i >= len(nums) {
				return
			}

			i--
		}
	}
}
