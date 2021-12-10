package main

func numTilings(n int) int {
	m := 1000000007
	dp := [1001]int{1: 1, 2: 2, 3: 5}
	if n <= 3 {
		return dp[n]
	}

	for i := 4; i <= n; i++ {
		dp[i] = 2*dp[i-1] + dp[i-3]
		dp[i] %= m
	}

	return dp[n]
}
