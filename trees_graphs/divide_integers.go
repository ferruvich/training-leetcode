package trees_graphs

import "math"

var (
	max = int(int32(math.Inf(1)) - 1)
	min = int(int32(math.Inf(-1)))
)

func divide(dividend int, divisor int) int {
	sign := 1
	if dividend < 0 {
		dividend = -dividend
		sign *= -1
	}
	if divisor < 0 {
		divisor = -divisor
		sign *= -1
	}

	res := computeDivisionWithBitShift(dividend, divisor)

	if res > max {
		if sign < 0 {
			return min
		}

		return max
	}

	return res * sign
}

func computeDivisionWithSubtractions(dividend int, divisor int) int {
	if divisor == 1 {
		return dividend
	}

	var count int
	for dividend >= divisor {
		dividend -= divisor
		count++
	}

	return count
}

func computeDivisionWithBitShift(dividend int, divisor int) int {
	highestDouble := divisor
	highestPowerOfTwo := 1
	for dividend >= highestDouble+highestDouble {
		highestPowerOfTwo += highestPowerOfTwo
		highestDouble += highestDouble
	}

	var quotient int
	for dividend >= divisor {
		if dividend >= highestDouble {
			quotient += highestPowerOfTwo
			dividend -= highestDouble
		}
		highestPowerOfTwo >>= 1
		highestDouble >>= 1
	}

	return quotient
}