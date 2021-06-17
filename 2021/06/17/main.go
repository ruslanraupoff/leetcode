package main

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	p, r, n := -1, 0, 0
	for i := 0; i < len(nums); i++ {
		if left <= nums[i] && nums[i] <= right {
			n = i - p
		}
		if nums[i] > right {
			n, p = 0, i
		}
		r += n
	}
	return r
}
