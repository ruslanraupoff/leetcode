package main

func longestOnes(nums []int, k int) int {
	l, r := 0, 0
	for r = range nums {
		k -= 1 ^ nums[r]
		if k < 0 {
			k += 1 ^ nums[l]
			l += 1
		}
	}
	return r - l + 1
}
