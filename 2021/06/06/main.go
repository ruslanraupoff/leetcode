package main

import "sort"

func longestConsecutive(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	sort.Ints(nums)
	mx, r := 0, 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i]-1 {
			r++
		} else if nums[i] != nums[i-1] {
			r = 1
		}
		if r > mx {
			mx = r
		}
	}
	return mx
}
