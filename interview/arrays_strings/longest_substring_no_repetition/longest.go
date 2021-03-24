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

const (
	uintSize = 32 << (^uint32(0) >> 32 & 1)

	maxInt  = 1<<(uintSize-1) - 1
	minInt  = -maxInt - 1
)

func myAtoi(s string) int {
	var numberStart, numberEnd int
	var startingZeros, numberStarted bool
	r := []rune(s)

	for i:=0; i<len(r); i++ {
		el := r[i]
		switch {
		case string(el) == " ":
			if numberStarted {
				numberEnd = i
				break
			}
			continue
		case el == 48: //0
			startingZeros = true
		case el == 43, el == 45: // +, -
			if startingZeros || numberStarted {return 0}
			numberStarted = true
			numberStart = i
		case el >= 49 && el <= 57:
			if !numberStarted {
				numberStarted = true
				numberStart = i
			}
		default:
			if !numberStarted {return 0}

			numberEnd = i
			break
		}
	}
	if numberEnd == 0 {
		numberEnd = len(r)
	}

	return atoi(string(r[numberStart:numberEnd]))
}

func atoi(s string) int {
	sign := 1
	multiplier := 1
	res := 0
	r := []rune(s)

	for i:=len(r)-1; i>=0; i-- {
		el := r[i]
		switch {
		case el == 45: // -
			sign = -1
		case el == 48: // 0
			multiplier *= 10
		case el >= 49 && el <= 57: // 1-9
			res += (int(el)-48) * multiplier
			multiplier *= 10
		}
	}

	if res < minInt {
		res = minInt
	}
	if res > maxInt {
		res = maxInt
	}

	return res*sign
}