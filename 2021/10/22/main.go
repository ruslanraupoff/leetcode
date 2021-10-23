package main

import (
	"sort"
)

func frequencySort(s string) string {
	r, mp := "", ['z' + 1]int{}
	for i := 0; i < len(s); i++ {
		mp[s[i]] += 1
	}
	var l []string
	for k := range mp {
		if mp[k] == 0 {
			continue
		}
		l = append(l, makeString(byte(k), mp[k]))
	}
	sort.Slice(l, func(i, j int) bool {
		return len(l[i]) > len(l[j])
	})
	for _, v := range l {
		r += v
	}
	return r
}

func makeString(b byte, n int) string {
	s := make([]byte, n)
	for i := range s {
		s[i] = b
	}
	return string(s)
}
