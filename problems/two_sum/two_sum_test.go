package two_sum_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/training-leetcode/two_sum"
)

func TestTwoSumBruteForce(t *testing.T) {
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
				res := two_sum.TwoSumBruteForce(tt.nums, tt.target)
				for _, e := range tt.out {
					assert.Contains(t, res, e)
				}
			},
		)
	}
}

func TestTwoSumMapListTraversedTwice(t *testing.T) {
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
				res := two_sum.TwoSumMapListTraversedTwice(tt.nums, tt.target)
				for _, e := range tt.out {
					assert.Contains(t, res, e)
				}
			},
		)
	}
}

func TestTwoSumMapSingleListTraverse(t *testing.T) {
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
				res := two_sum.TwoSumMapSingleListTraverse(tt.nums, tt.target)
				for _, e := range tt.out {
					assert.Contains(t, res, e)
				}
			},
		)
	}
}
