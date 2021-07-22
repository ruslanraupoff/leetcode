package main

func partitionDisjoint(nums []int) int {
	p, n, mx := 0, nums[0], nums[0]
	for i := 0; i < len(nums); i++ {
		if mx < nums[i] {
			mx = nums[i]
		}
		if nums[i] < n {
			p = i
			n = mx
		}
	}
	return p + 1
}
