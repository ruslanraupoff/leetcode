package main

func calculateMinimumHP(dungeon [][]int) int {
	n, m := len(dungeon[0]), len(dungeon)
	if m == 0 || n == 0 {
		return 1
	}

	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = 1<<63 - 1
		}
	}

	dp[m][n-1] = 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mn := min(dp[i+1][j], dp[i][j+1]) - dungeon[i][j]
			dp[i][j] = max(mn, 1)
		}
	}

	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
