package main

func findPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	dp := [50][50]int{}

	for k := 0; k < maxMove; k++ {
		p := make([]int, n)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				r := 0

				if i == 0 {
					r++
				} else {
					r += p[j]
				}

				if j == 0 {
					r++
				} else {
					r += p[j-1]
				}

				if i == m-1 {
					r++
				} else {
					r += dp[i+1][j]
				}

				if j == n-1 {
					r++
				} else {
					r += dp[i][j+1]
				}

				r %= (1e9 + 7)
				p[j] = dp[i][j]
				dp[i][j] = r
			}
		}
	}

	return dp[startRow][startColumn]
}
