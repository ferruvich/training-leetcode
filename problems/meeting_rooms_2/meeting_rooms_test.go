package meeting_rooms_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	meeting_rooms "github.com/ferruvich/training-leetcode/problems/meeting_rooms_2"
)

func TestMinMeetingRooms(t *testing.T) {
	for _, tt := range []struct {
		input         [][]int
		expectedRooms int
	}{
		{input: [][]int{{0, 30}, {5, 10}, {15, 20}}, expectedRooms: 2},
		{input: [][]int{{0, 30}, {5, 10}, {15, 20}, {30, 40}, {30, 50}, {20, 25}, {25, 31}}, expectedRooms: 3},
		{input: [][]int{{7, 10}, {2, 4}}, expectedRooms: 1},
		{input: [][]int{{7, 10}, {2, 4}, {1, 2}}, expectedRooms: 1},
	} {
		t.Run(
			fmt.Sprintf("should return rooms=%d as input is %v", tt.expectedRooms, tt.input),
			func(t *testing.T) {
				res := meeting_rooms.MinMeetingRooms(tt.input)
				assert.Equal(t, tt.expectedRooms, res)
			},
		)
	}
}
