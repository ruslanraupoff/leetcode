package main

func numSquares(n int) int {
	p := []int{}
	for i := 1; i*i <= n; i++ {
		p = append(p, i*i)
	}

	dp := make([]int, n+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = (n + 1) * (n + 1)
	}

	for _, v := range p {
		for i := v; i < len(dp); i++ {
			if dp[i] > dp[i-v]+1 {
				dp[i] = dp[i-v] + 1
			}
		}
	}

	return dp[n]
}
