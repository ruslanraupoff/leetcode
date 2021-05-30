package main

import "sort"

func maximumGap(nums []int) int {
	sort.Ints(nums)
	mx := 0
	for i := 1; i < len(nums); i++ {
		d := nums[i] - nums[i-1]
		if mx < d {
			mx = d
		}
	}
	return mx
}
