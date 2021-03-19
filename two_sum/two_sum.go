package two_sum

// This has
// O(n^2) time complexity, as we traverse the remaining list once for each element
// O(1) space complexity, as we're not saving anything but the input and the output in memory
func TwoSumBruteForce(nums []int, target int) []int {
	for i, n := range nums {
		for j, m := range nums {
			if n+m > target || i == j {
				continue
			}

			if m+n == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

// This has
// O(2n)->O(n) time complexity, as we traverse the list twice
// O(n) space complexity, as we're using a map, saving each value as key
func TwoSumMapListTraversedTwice(nums []int, target int) []int {
	elements := make(map[int]int, len(nums))
	for i, n := range nums {
		elements[n] = i
	}

	for i, n := range nums {
		complement := target - n
		if j, ok := elements[complement]; ok && i != j{
			return []int{i,j}
		}
	}

	return nil
}

func TwoSumMapSingleListTraverse(nums []int, target int) []int {
	elements := make(map[int]int, len(nums))
	for i, n := range nums {
		complement := target - n
		if j, ok := elements[complement]; ok && i != j{
			return []int{i,j}
		}

		elements[n] = i
	}

	return nil
}