package two_sum_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/training-leetcode/two_sum"
)

func TestTwoSum(t *testing.T) {
	for _, tt := range []struct {
		nums   []int
		target int
		out    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{-3, 4, 3, 90}, 0, []int{0, 2}},
	} {
		t.Run(
			fmt.Sprintf("should return %v having nums=%v and targed=%d", tt.out, tt.nums, tt.target),
			func(t *testing.T) {
				resBruteForce := two_sum.TwoSumBruteForce(tt.nums, tt.target)
				resMapListTraversedTwice := TwoSumMapListTraversedTwice(tt.nums, tt.target)
				resSingleListTraverse := TwoSumMapSingleListTraverse(tt.nums, tt.target)
				assert.Equal(t, resBruteForce, tt.out)
				assert.Equal(t, resMapListTraversedTwice, tt.out)
				assert.Equal(t, resSingleListTraverse, tt.out)
			},
		)
	}
}
