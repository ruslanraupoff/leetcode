package main

func cherryPickup(grid [][]int) int {
	var dp [71][71][71]int
	m, n := len(grid), len(grid[0])

	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				r := grid[i][j]
				if j != k {
					r += grid[i][k]
				}
				if i != m-1 {
					mx := 0
					for x := j - 1; x <= j+1; x++ {
						for y := k - 1; y <= k+1; y++ {
							if x >= 0 && x < n && y >= 0 && y < n {
								mx = max(mx, dp[i+1][x][y])
							}
						}
					}
					r += mx
				}

				dp[i][j][k] = r
			}
		}
	}

	return dp[0][0][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
