package main

import (
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	d, ln := 100000, len(nums)
	sort.Ints(nums)
	for i := 0; i < ln; i++ {
		l, r := i+1, ln-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if abs(target-sum) < abs(d) {
				d = target - sum
			}
			if sum < target {
				l += 1
			} else {
				r -= 1
			}
		}
		if d == 0 {
			break
		}
	}
	return target - d
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
