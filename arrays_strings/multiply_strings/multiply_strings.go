package multiply_strings

func multiply(num1 string, num2 string) string {
	results := make([]int, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			res := int((num1[i] - '0') * (num2[j] - '0'))
			sum := results[i+j+1] + int(res)

			results[i+j+1] = sum % 10
			results[i+j] += sum / 10
		}
	}

	var resultsString []rune
	for _, el := range results {
		if len(resultsString) != 0 || el != 0 {
			resultsString = append(resultsString, rune(el)+'0')
		}
	}

	if len(resultsString) == 0 {
		return "0"
	}

	return string(resultsString)
}
