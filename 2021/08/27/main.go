package main

import "sort"

func findLUSlength(strs []string) int {
	count := make(map[string]int, len(strs))
	for _, s := range strs {
		count[s]++
	}
	strs = strs[:0]
	for s := range count {
		strs = append(strs, s)
	}

	sort.Sort(stringSlice(strs))

	for i, s := range strs {
		if count[s] > 1 {
			continue
		}
		if !isSubOf(s, strs[:i]) {
			return len(s)
		}
	}

	return -1
}

func isSubOf(s string, ss []string) bool {
	for i := range ss {
		if isSub(s, ss[i]) {
			return true
		}
	}
	return false
}

func isSub(a, b string) bool {
	aLen, bLen := len(a), len(b)
	i, j := 0, 0

	for i < aLen && j < bLen {
		if a[i] == b[j] {
			i++
		}
		j++
	}

	return i == aLen
}

type stringSlice []string

func (ss stringSlice) Len() int { return len(ss) }

func (ss stringSlice) Less(i, j int) bool { return len(ss[i]) > len(ss[j]) }

func (ss stringSlice) Swap(i, j int) { ss[i], ss[j] = ss[j], ss[i] }
