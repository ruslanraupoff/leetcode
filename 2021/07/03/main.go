package main

import "math"

func maxSumSubmatrix(matrix [][]int, k int) int {
	n, m := len(matrix), len(matrix[0])
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := 0; j <= m; j++ {
			dp[i][j] = 0
		}
	}

	r := math.MinInt32
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = matrix[i-1][j-1] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
			if dp[i][j] == k {
				return k
			}
			for p := 1; p <= i; p++ {
				for q := 1; q <= j; q++ {
					s := dp[i][j] - dp[p-1][j] - dp[i][q-1] + dp[p-1][q-1]
					if s <= k && s > r {
						r = s
					}
				}
			}
			if r == k {
				return k
			}
		}
	}
	return r
}
