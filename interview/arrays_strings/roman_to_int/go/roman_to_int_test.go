package roman_to_int_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/training-leetcode/interview/arrays_strings/roman_to_int/go"
)

func TestRomanToInt(t *testing.T) {
	for _, tt := range []struct {
		inputString   string
		expectedValue int
	}{
		{inputString: "III", expectedValue: 3},
		{inputString: "XX", expectedValue: 20},
		{inputString: "IV", expectedValue: 4},
		{inputString: "IX", expectedValue: 9},
		{inputString: "LVIII", expectedValue: 58},
		{inputString: "MCMXCIV", expectedValue: 1994},
	} {
		t.Run(
			fmt.Sprintf("should return correctly return %d as the string is %s", tt.expectedValue, tt.inputString),
			func(t *testing.T) {
				res := roman_to_int.RomanToInt(tt.inputString)
				assert.Equal(t, tt.expectedValue, res)
			},
		)
	}
}
