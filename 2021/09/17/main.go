package main

func intersect(nums1 []int, nums2 []int) []int {
	res := []int{}
	m1 := getInts(nums1)
	m2 := getInts(nums2)

	if len(m1) > len(m2) {
		m1, m2 = m2, m1
	}
	for n := range m1 {
		m1[n] = min(m1[n], m2[n])
	}

	for n, size := range m1 {
		for i := 0; i < size; i++ {
			res = append(res, n)
		}
	}

	return res
}

func getInts(a []int) map[int]int {
	res := make(map[int]int, len(a))

	for i := range a {
		res[a[i]]++
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
