package valid_parenthesis

func ValidParenthesis(s string) bool {
	var toBeClosed []string
	openClosed := map[string]string{"(": ")", "[": "]", "{": "}"}

	for _, c := range []rune(s) {
		switch sc := string(c); sc {
		case "(", "[", "{":
			toBeClosed = append(toBeClosed, sc)
		default:
			if len(toBeClosed) == 0 {
				return false
			}

			o := toBeClosed[len(toBeClosed)-1]
			if sc != openClosed[o] {
				return false
			}
			toBeClosed = toBeClosed[:len(toBeClosed)-1]
		}
	}

	return len(toBeClosed) == 0
}
