package main

func findMaxConsecutiveOnes(nums []int) int {
	r, i, cnt := 0, 0, 0
	for i < len(nums) {

		for i < len(nums) && nums[i] == 1 {
			i += 1
			cnt += 1
		}
		r = max(r, cnt)
		i += 1
		cnt = 0
	}
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
