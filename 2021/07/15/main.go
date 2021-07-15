package main

import "sort"

func triangleNumber(nums []int) int {
	c, n := 0, len(nums)
	sort.Ints(nums)
	for i := n - 1; i > -1; i-- {
		l, r := 0, i-1
		for l < r {
			if nums[l]+nums[r] > nums[i] {
				c += r - l
				r -= 1
			} else {
				l += 1
			}
		}
	}
	return c
}
