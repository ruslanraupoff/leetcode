package main

func kInversePairs(n int, k int) int {
	m, dp := 1000000007, make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
		dp[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			if j >= i {
				dp[i][j] = dp[i][j-1] + dp[i-1][j] - dp[i-1][j-i]
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			}
			dp[i][j] %= m
			if dp[i][j] < 0 {
				dp[i][j] += m
			}
		}
	}
	return dp[n][k]
}
