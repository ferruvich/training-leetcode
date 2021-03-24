package alien_dictionary

import (
	"sort"
	"strings"
)

var lexOrder = map[string]int{}

func IsAlienSorted(words []string, order string) bool {
	for i, c := range strings.Split(order, "") {
		lexOrder[c] = i
	}


	return sort.SliceIsSorted(words, func(x, y int) bool {
		return isLess(words[x], words[y])
	})
}

func isLess(first, second string) bool {
	if len(first) == 0 {
		if len(second) == 0 {
			// is equal, not less
			return false
		}

		return true
	}
	if len(second) == 0 {
		return false
	}

	firstRunes, secondRunes := []rune(first), []rune(second)
	firstLetterEqual := lexOrder[string(firstRunes[0])] == lexOrder[string(secondRunes[0])]
	if !firstLetterEqual {
		return lexOrder[string(firstRunes[0])] < lexOrder[string(secondRunes[0])]
	}

	return isLess(string(firstRunes[1:]), string(secondRunes[1:]))
}
