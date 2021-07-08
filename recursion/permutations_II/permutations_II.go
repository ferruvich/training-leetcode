package permutations_II

var (
	results [][]int
)

func permuteUnique(nums []int) [][]int {
	counter := map[int]int{}
	results = [][]int{}

	for _, el := range nums {
		counter[el] += 1
	}

	backTrack(nums, []int{}, counter)

	return results
}

func backTrack(nums, combination []int, counter map[int]int) {
	if len(combination) == len(nums) {
		cpy := make([]int, len(combination))
		copy(cpy, combination)

		results = append(results, cpy)
		return
	}

	for num, numCounter := range counter {
		if numCounter > 0 {
			// Adding number to combination, reducing its occurrences
			combination = append(combination, num)
			counter[num] -= 1
			// continuing with this combination
			backTrack(nums, combination, counter)
			// Setting everything back to previous status, as we're
			// backtracking to next solution
			combination = combination[:len(combination)-1]
			counter[num] += 1
		}
	}
}
