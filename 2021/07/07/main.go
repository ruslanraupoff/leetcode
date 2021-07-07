package main

import "sort"

func kthSmallest(matrix [][]int, k int) int {
	l := []int{-10000000001}
	for _, m := range matrix {
		l = append(l, m...)
	}
	sort.Ints(l)
	return l[k]
}
