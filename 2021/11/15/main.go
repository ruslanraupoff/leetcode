package main

import "sort"

func largestDivisibleSubset(nums []int) []int {
	size := len(nums)
	if size == 0 {
		return []int{}
	}

	sort.Ints(nums)
	mx := 1
	idx := 0

	dp := make([]int, size)
	for i := range dp {
		dp[i] = 1
	}

	next := make([]int, size)

	for i := size - 2; 0 <= i; i-- {
		for j := size - 1; i < j; j-- {
			if nums[j]%nums[i] != 0 {
				continue
			}
			if dp[i] < dp[j]+1 {
				next[i] = j
				dp[i] = dp[j] + 1
			}
			if mx < dp[i] {
				mx = dp[i]
				idx = i
			}
		}
	}

	res := make([]int, mx)
	for i := range res {
		res[i] = nums[idx]
		idx = next[idx]
	}

	return res
}
