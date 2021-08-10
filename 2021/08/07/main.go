package main

func minCut(s string) int {
	n := len(s)
	l := make([]int, n+1)
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, n)
		for j := 0; j <= n; j++ {
			dp[i][j] = false
		}
	}

	for j := 0; j < n; j++ {
		for i := j; i > -1; i-- {
			if s[i] == s[j] && (j-i < 2 || dp[i+1][j]) {
				dp[i][j] = true
				if l[j] > l[i]+1 {
					l[j] = l[i] + 1
				}
			}
		}
	}

	return l[0]
}
