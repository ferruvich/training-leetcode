package longest_substring_no_repetition

import (
	"strings"
)

// O(n) time, as we are traversing the substring once
// O(1) space, as we don't store anything
func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}

	i, j, maxLen := 0, 1, 0
	r := []rune(s)
	for {
		if j == len(s) {
			if maxLen == 0 {
				return len(s)
			}
			return maxLen
		}
		if strings.ContainsRune(string(r[i:j]), r[j]) {
			i++
		} else {
			j++
		}

		if len(r[i:j]) > maxLen {
			maxLen = len(r[i:j])
		}
	}
}
