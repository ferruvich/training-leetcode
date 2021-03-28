package merge_intervals_test

import (
	"fmt"
	"testing"

	"github.com/ferruvich/training-leetcode/arrays_strings/merge_intervals"
	"github.com/stretchr/testify/assert"
)

func TestMergeIntervals(t *testing.T) {
	for _, tt := range []struct {
		intervals [][]int
		result    [][]int
	}{
		{
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			result:    [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			intervals: [][]int{{1, 4}, {4, 5}},
			result:    [][]int{{1, 5}},
		},
	} {
		t.Run(
			fmt.Sprintf("should return correctly intervals %v when initial are %v", tt.result, tt.intervals),
			func(t *testing.T) {
				res := merge_intervals.MergeIntervals(tt.intervals)
				assert.Equal(t, tt.result, res)
			},
		)
	}
}
