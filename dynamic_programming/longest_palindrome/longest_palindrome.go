package longest_palindrome

func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}
	start, end := 0, 0

	for i := range s {
		lenOdd := expandAroundCenter(s, i, i)
		lenEven := expandAroundCenter(s, i, i+1)

		maxLen := lenOdd
		if maxLen < lenEven {
			maxLen = lenEven
		}

		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) int {
	l, r := left, right

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return r - l - 1
}
