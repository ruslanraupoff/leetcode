package main

func minMoves2(nums []int) int {
	r := 10000000001
	for i := 0; i < len(nums); i++ {
		s := 0
		for j := 0; j < len(nums); j++ {
			if nums[i] > nums[j] {
				s += nums[i] - nums[j]
			} else {
				s += nums[j] - nums[i]
			}
		}
		if s < r {
			r = s
		}
	}
	return r
}
