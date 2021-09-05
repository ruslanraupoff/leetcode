package main

import "sort"

func orderlyQueue(s string, k int) string {
	n := len(s)
	if n == 1 {
		return s
	}
	if k > 1 {
		return sorted(s)
	}
	s2 := s + s
	r := "z"
	for i := 1; i < n; i++ {
		r += "z"
	}
	for i := 0; i < n; i++ {
		t := s2[i : i+n]
		if t < r {
			r = t
		}
	}
	return r
}

func sorted(s string) string {
	bytes := []byte(s)
	sort.Slice(bytes, func(i int, j int) bool {
		return bytes[i] < bytes[j]
	})
	return string(bytes)
}
