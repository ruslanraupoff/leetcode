package main

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return 0
	}

	mxe := 0
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
				mxe = max(mxe, dp[i][j])
			}
		}
	}

	return mxe * mxe
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
