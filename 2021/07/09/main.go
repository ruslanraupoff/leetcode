package main

func lengthOfLIS(nums []int) int {
	dp = make([][]int)
	for i := 0; i < len(nums); i++ {
		idx := bs(dp, nums[i])
		k = len(dp)
		if idx == k {
			arr = []int{}
			if k != 0 {
				for j := 0; j < dp[k-1]; j++ {
					arr = append(arr, j)
				}
			}
			arr = append(arr, nums[i])
			dp = append(dp, arr)
		} else if dp[idx][-1] > nums[i] {
			dp[idx][-1] = nums[i]
		}
	}
	return dp[-1]
}

func bs(dp []int, t int) int {
	l, r := 0, len(dp)-1
	for l <= r {
		m := l + (r-l)/2
		d := dp[m]
		if d[-1] == t {
			return m
		} else if d[-1] < t {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}
