package atoi_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/training-leetcode/interview/arrays_strings/atoi"
)

func TestMyAtoi(t *testing.T) {
	for _, tt := range []struct {
		inputString    string
		expectedNumber int
	}{
		{inputString: "", expectedNumber: 0},
		{inputString: "42", expectedNumber: 42},
		{inputString: "   -42", expectedNumber: -42},
		{inputString: "4193 with words", expectedNumber: 4193},
		{inputString: "words and 987", expectedNumber: 0},
		{inputString: "-91283472332", expectedNumber: -2147483648},
	} {
		t.Run(
			fmt.Sprintf("should return %d as the string is \"%s\"", tt.expectedNumber, tt.inputString),
			func(t *testing.T) {
				atoied := atoi.MyAtoi(tt.inputString)
				assert.Equal(t, tt.expectedNumber, atoied)
			},
		)
	}
}
