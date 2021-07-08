package main

func findLength(nums1 []int, nums2 []int) int {
	r, n, m := 0, len(nums1), len(nums2)

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if r < dp[i][j] {
					r = dp[i][j]
				}
			}
		}
	}

	return r
}
