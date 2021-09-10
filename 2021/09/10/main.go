package main

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	dp := make([]map[int]int, n)
	res := 0
	for i := 1; i < n; i++ {
		dp[i] = make(map[int]int)
		for j := 0; j < i; j++ {
			d := nums[i] - nums[j]
			dp[i][d] += dp[j][d] + 1
			if dp[j][d] > 0 {
				res += dp[j][d]
			}
		}
	}
	return res
}
