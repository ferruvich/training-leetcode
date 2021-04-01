package group_anagrams_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ferruvich/training-leetcode/arrays_strings/group_anagrams"
)

func TestGroupAnagrams(t *testing.T) {
	for _, tt := range []struct {
		input          []string
		expectedOutput [][]string
	}{
		{input: []string{}, expectedOutput: [][]string{}},
		{input: []string{"a", "b", "a"}, expectedOutput: [][]string{{"a", "a"}, {"b"}}},
		{input: []string{"abab", "baba", "bababb"}, expectedOutput: [][]string{{"abab", "baba"}, {"bababb"}}},
		{
			input:          []string{"ska", "tac", "tesco", "cost", "cats", "scat", "skip", "ask", "cast", "cat"},
			expectedOutput: [][]string{{"cat", "tac"}, {"tesco"}, {"cost"}, {"cats", "cast", "scat"}, {"skip"}, {"ask", "ska"}},
		},
	} {
		t.Run(
			fmt.Sprintf("should correctly return %v as the list of values is %v", tt.expectedOutput, tt.input),
			func(t *testing.T) {
				anagrams := group_anagrams.GroupAnagrams(tt.input)
				elementsNo := 0
				for _, an := range anagrams {
					for _, s := range an {
						elementsNo++
						assert.Contains(t, tt.input, s)
					}
				}
				assert.Equal(t, len(tt.input), elementsNo)
			},
		)
	}
}
