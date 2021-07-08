package letter_combination_phone_number

var m = map[rune][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func LetterCombinations(digits string) []string {
	var res []string

	if len(digits) == 0 {
		return nil
	}

	if len(digits) == 1 {
		r := []rune(digits)
		return m[r[0]]
	}

	comb := LetterCombinations(digits[1:])
	for _, c := range comb {
		for _, l := range m[rune(digits[0])] {
			res = append(res, l+c)
		}
	}

	return res
}
