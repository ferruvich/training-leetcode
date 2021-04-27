package search_rotated_sorted_array

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	rotationBase := findRotationBase(nums)

	if rotationBase == 0 {
		return binarySearch(nums, target, 0, len(nums)-1)
	}

	if target >= nums[0] {
		return binarySearch(nums, target, 0, rotationBase)
	}

	return binarySearch(nums, target, rotationBase, len(nums)-1)
}

func findRotationBase(nums []int) int {
	left, right := 0, len(nums)-1

	pivot := 0
	for left <= right {
		pivot = (left + right) / 2

		if pivot+1 >= len(nums) {
			return 0
		}

		if nums[pivot] > nums[pivot+1] {
			return pivot + 1
		}

		if nums[pivot] < nums[left] {
			right = pivot - 1
		} else {
			left = pivot + 1
		}
	}

	return 0
}

func binarySearch(nums []int, target, left, right int) int {
	if left > right {
		return -1
	}

	pivot := (left + right) / 2
	if nums[pivot] == target {
		return pivot
	}

	if target < nums[pivot] {
		return binarySearch(nums, target, left, pivot-1)
	}

	return binarySearch(nums, target, pivot+1, right)
}
