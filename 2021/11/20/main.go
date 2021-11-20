package main

func singleNonDuplicate(nums []int) int {
	n := len(nums) - 1
	l, r := 0, n-1

	for l < r {
		m := l + (r-l)/2
		if m%2 == 1 {
			m--
		}

		if nums[m] != nums[m+1] {
			r = m
		} else {
			l = m + 2
		}
	}

	return nums[l]
}
