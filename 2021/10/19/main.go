package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	l := make(map[int]int)
	for i, n := range nums2 {
		l[n] = i
	}

	r := make([]int, len(nums1))
	for i, n := range nums1 {
		r[i] = -1
		for j := l[n] + 1; j < len(nums2); j++ {
			if n < nums2[j] {
				r[i] = nums2[j]
				break
			}
		}
	}

	return r
}
