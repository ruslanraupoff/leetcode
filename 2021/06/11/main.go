package main

func stoneGameVII(stones []int) int {
	d := len(stones)
	dp := make([][]int, d)
	for i := range dp {
		dp[i] = make([]int, d)
	}
	for i := d - 1; i > -1; i-- {
		sm := stones[i]
		for j := i + 1; j < d; j++ {
			sm += stones[j]
			dp[i][j] = max(sm-stones[i]-dp[i+1][j], sm-stones[j]-dp[i][j-1])
		}
	}
	return dp[0][d-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
