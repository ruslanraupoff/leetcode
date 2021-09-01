package main

func arrayNesting(nums []int) int {
	r := 0
	for i := range nums {
		t := 1
		for nums[i] != i {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
			t++
		}
		if r < t {
			r = t
		}
	}
	return r
}
