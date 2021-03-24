package islands_number_test

import (
	"fmt"
	"testing"

	"github.com/ferruvich/training-leetcode/islands_number"
	"github.com/stretchr/testify/assert"
)

func TestNumIslands(t *testing.T) {
	for _, tt := range []struct {
		grid          [][]byte
		islandsNumber int
	}{
		{
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			islandsNumber: 1,
		},
		{
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			islandsNumber: 3,
		},
	} {
		t.Run(
			fmt.Sprintf("should return correctly spot %d islands with %v grid", tt.islandsNumber, tt.grid),
			func(t *testing.T) {
				n := islands_number.NumIslands(tt.grid)
				assert.Equal(t, tt.islandsNumber, n)
			},
		)
	}
}
