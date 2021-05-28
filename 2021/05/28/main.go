package main

func maximumUniqueSubarray(nums []int) int {
	l, s, r, i := [10001]bool{}, 0, 0, 0
	for _, n := range nums {
		for l[n] {
			l[nums[i]] = false
			s -= nums[i]
			i++
		}
		l[n] = true
		s += n
		if s > r {
			r = s
		}
	}
	return r
}
