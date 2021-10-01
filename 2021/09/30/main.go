package main

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	s, n := 0, len(nums)
	mx := nums[0]
	for _, v := range nums {
		s += v
		if mx < v {
			mx = v
		}
	}

	if s%k != 0 || s < mx*k {
		return false
	}

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	u := make([]bool, n)

	return dfs(nums, u, k, 0, 0, s/k)
}

func dfs(nums []int, u []bool, k, j, s, t int) bool {
	if k == 1 {
		return true
	}

	if s == t {
		return dfs(nums, u, k-1, 0, 0, t)
	}

	for i := j; i < len(nums); i++ {
		if !u[i] && s+nums[i] <= t {
			u[i] = true
			if dfs(nums, u, k, i+1, s+nums[i], t) {
				return true
			}
			u[i] = false
		}
	}

	return false
}
