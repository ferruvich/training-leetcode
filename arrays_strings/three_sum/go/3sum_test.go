package three_sum_test

import (
	"fmt"
	three_sum "github.com/ferruvich/training-leetcode/arrays_strings/three_sum/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum(t *testing.T) {
	for _, tt := range []struct {
		input []int
		out   [][]int
	}{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{2, -1, -1}, {1, 0, -1}}},
		{[]int{}, [][]int{}},
		{[]int{0, 1}, [][]int{}},
		{[]int{0}, [][]int{}},
		{[]int{1, 2, 3, 4, 5, 6}, [][]int{}},
		{[]int{-1, 0, -1, 1, 2, 2, -4}, [][]int{{1, 0, -1}, {2, 2, -4}, {2, -1, -1}}},
	} {
		t.Run(fmt.Sprintf("should return %v as input is %v", tt.out, tt.input), func(t *testing.T) {
			res := three_sum.ThreeSum(tt.input)
			for _, el := range tt.out {
				assert.Contains(t, res, el)
			}
		})
	}
}
