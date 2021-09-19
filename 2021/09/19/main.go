package main

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)

	p, dp := 0, make([]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = 1
	}
	for j := 0; j < n; j++ {
		dp[j], p = 0, dp[j]
		for i := j + 1; i < m+1; i++ {
			if t[j] == s[i-1] {
				dp[i], p = dp[i-1]+p, dp[i]
			} else {
				dp[i], p = dp[i-1], dp[i]
			}
		}
	}

	return dp[m]
}
