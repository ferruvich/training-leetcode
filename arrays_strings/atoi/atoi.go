package atoi

const (
	uintSize = 32 << (^uint32(0) >> 32 & 1)

	maxInt = 1<<(uintSize-1) - 1
	minInt = -maxInt - 1
)

func MyAtoi(s string) int {
	base := 10
	res := 0
	sign := 0
	numberStarted := false
	numberStartingAt := -1

	for i, r := range []rune(s) {
		switch {
		case r == 32: //space
			if numberStarted {
				return clampToInt32(res * sign)
			}
			if sign != 0 {
				return 0
			}
		case r == 43: // +
			if numberStarted {
				return clampToInt32(res * sign)
			}
			if sign != 0 {
				return 0
			}
			sign = 1
		case r == 45: // -
			if numberStarted {
				return clampToInt32(res * sign)
			}
			if sign != 0 {
				return 0
			}
			sign = -1
		case r == 48: // 0
			numberStarted = true
			if numberStarted {
				res *= base
			}
		case r >= 49 && r <= 57: // 0-9
			if numberStartingAt == -1 {
				numberStartingAt = i
			}
			if sign == 0 {
				sign = 1
			}
			numberStarted = true
			res *= base
			res += int(r) - 48

			if i-numberStartingAt > 11 {
				if sign == -1 {
					return minInt
				}
				return maxInt
			}
		default:
			if numberStarted {
				return clampToInt32(res * sign)
			}
			return 0
		}
	}

	return clampToInt32(res * sign)
}

func clampToInt32(val int) int {
	if val < minInt {
		return minInt
	}
	if val > maxInt {
		return maxInt
	}
	return val
}
