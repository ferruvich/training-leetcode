package decode_ways_test

import (
	"fmt"
	"testing"

	"github.com/ferruvich/training-leetcode/arrays_strings/decode_ways"
	"github.com/stretchr/testify/assert"
)

func TestNumIslands(t *testing.T) {
	for _, tt := range []struct {
		input          string
		expectedResult int
	}{
		{input: "12", expectedResult: 2},
		{input: "226", expectedResult: 3},
		{input: "06", expectedResult: 0},
		{input: "0", expectedResult: 0},
		{input: "111111111111111111111111111111111111111111111", expectedResult: 1836311903},
	} {
		t.Run(
			fmt.Sprintf("should return %d as the string is %s", tt.expectedResult, tt.input),
			func(t *testing.T) {
				res := decode_ways.NumDecodings(tt.input)
				assert.Equal(t, tt.expectedResult, res)
			},
		)
	}
}
