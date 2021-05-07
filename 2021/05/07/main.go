package main

func minDistance(word1, word2 string) int {
	l1, l2 := len(word1), len(word2)

	dp := make([][]int, l1+1)
	for i := range dp {
		dp[i] = make([]int, l2+1)
	}

	for i := 1; i <= l1; i++ {
		for j := 1; j <= l2; j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
		}
	}

	return l1 + l2 - dp[l1][l2]*2
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
