package main

func canPartition(nums []int) bool {
	s := 0
	for _, n := range nums {
		s += n
	}

	if s&1 == 1 {
		return false
	}

	s = s >> 1
	n := len(nums)

	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, s+1)
	}

	for i := 0; i < n+1; i++ {
		dp[i][0] = true
	}

	for j := 1; j < s+1; j++ {
		dp[0][j] = false
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < s+1; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= nums[i-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[n][s]
}
