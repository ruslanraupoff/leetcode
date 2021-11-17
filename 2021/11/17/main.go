package main

func uniquePaths(m int, n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = 1
	}

	for i := 2; i <= m; i++ {
		for j := 2; j <= n; j++ {
			dp[j] += dp[j-1]
		}
	}

	return dp[n]
}
