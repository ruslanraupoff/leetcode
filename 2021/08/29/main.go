package main

func minPatches(nums []int, n int) int {
	i, mx, r := 0, 1, 0
	for mx <= n {
		if i < len(nums) && nums[i] <= mx {
			mx += nums[i]
			i++
		} else {
			mx <<= 1
			r++
		}
	}
	return r
}
