package roman_to_int

var romanToIntMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func RomanToInt(s string) int {
	r := []byte(s)
	res := 0
	for i:=len(r)-1; i>=0; i-- {
		n := romanToIntMap[string(r[i])]
		for j:=i-1; j>=0; j-- {
			m := romanToIntMap[string(r[j])]
			if m < n {
				n -= m
				i--
			}
		}
		res += n
	}

	return res
}
