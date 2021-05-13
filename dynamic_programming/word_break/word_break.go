package word_break

func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool, len(wordDict))
	for _, el := range wordDict {
		m[el] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && m[string(s[j:i])] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
