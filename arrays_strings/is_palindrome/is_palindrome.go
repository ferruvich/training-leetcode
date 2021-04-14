package is_palindrome

import "strings"

func IsPalindrome(s string) bool {
	runeS := []rune(strings.ToLower(s))
	j := len(runeS) - 1

	for i := 0; i < len(runeS); i++ {
		if needsCheck(runeS[i]) {
			if needsCheck(runeS[j]) {
				if runeS[i] != runeS[j] {
					return false
				}
			} else {
				i--
			}
			j--
		}
	}

	return true
}

func needsCheck(r rune) bool {
	return (r >= '0' && r <= '9') || (r >= 'a' && r <= 'z')
}
