package main

func findPeakElement(nums []int) int {
	l, r := -1, len(nums)
	m := l + (r-l)/2

	for l+1 <= m && m < r-1 {
		if nums[m] < nums[m+1] {
			l = m
		} else {
			r = m + 1
		}
		m = l + (r-l)/2
	}

	return l + 1
}
