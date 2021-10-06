package main

func findDuplicates(nums []int) []int {
	l, r := make(map[int]bool), make([]int, 0)
	for _, n := range nums {
		if _, ok := l[n]; ok {
			r = append(r, n)
		} else {
			l[n] = true
		}
	}
	return r
}
