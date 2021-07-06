package main

import "sort"

func minSetSize(arr []int) int {
	l, s := []int{}, make(map[int]int)
	for _, a := range arr {
		s[a]++
	}
	for _, v := range s {
		l = append(l, v)
	}
	sort.Slice(l, func(i, j int) bool {
		return l[i] > l[j]
	})
	c := 0
	for i, p := range l {
		c += p
		if c*2 >= len(arr) {
			return i + 1
		}
	}
	return 0
}
