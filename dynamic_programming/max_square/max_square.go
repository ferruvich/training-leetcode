package max_square

func maximalSquare(matrix [][]byte) int {
	row, col := len(matrix), len(matrix[0])
	maxValue := 0

	dp := make([][]int, row)

	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
		for j := 0; j < col; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					min := dp[i-1][j]
					if dp[i-1][j-1] < min {
						min = dp[i-1][j-1]
					}
					if dp[i][j-1] < min {
						min = dp[i][j-1]
					}

					dp[i][j] = min + 1
				}

				if dp[i][j] > maxValue {
					maxValue = dp[i][j]
				}
			}
		}
	}

	return maxValue * maxValue
}
