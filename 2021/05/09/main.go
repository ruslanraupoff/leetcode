package main

import "sort"

func isPossible(target []int) bool {
	sort.Ints(target)
	s, l := 0, len(target)
	for i := 0; i < l-1; i++ {
		s += target[i]
	}
	a := target[l-1]
	if l == 1 {
		return a == 1
	}
	if a == 1 || s == 1 {
		return true
	}
	if a < s || a%s == 0 {
		return false
	}
	a %= s
	s += a
	target[l-1] = a
	return isPossible(target)
}
