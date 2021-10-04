package main

func canJump(nums []int) bool {
	r := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i+nums[i] >= r {
			r = i
		}
	}

	return r == 0
}
