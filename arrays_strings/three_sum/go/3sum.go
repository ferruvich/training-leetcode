package three_sum

import "sort"

func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int

	for i:=0; i<len(nums) && nums[i] <= 0; i++ {
		if i==0 || nums[i-1] != nums[i] {
			sums := twoSum(nums[i+1:], 0-nums[i])
			for _, sum := range sums {
				res = append(res, append(sum, nums[i]))
			}
		}
	}
	return res
}

func twoSum(nums []int, target int) [][]int {
	m := map[int]int{}
	var res [][]int

	for j:=0; j< len(nums); j++ {
		el := nums[j]
		complement := target - el

		if k, ok := m[complement]; ok && k != j {
			res = append(res, []int{el, complement})
			for j+1 < len(nums) && nums[j] == nums[j+1] {
				j++
			}
		}
		m[el] = j
	}

	return res
}