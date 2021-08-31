package main

func findMin(nums []int) int {
	i, n := 1, len(nums)
	for i < n && nums[i-1] < nums[i] {
		i++
	}

	return nums[i%n]
}
