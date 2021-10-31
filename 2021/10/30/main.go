package main

func longestDupSubstring(s string) string {
	res := ""
	isExist := func(n int) bool {
		seen := make(map[string]bool, len(s)-n+2)
		for i := 0; i+n <= len(s); i++ {
			sub := s[i : i+n]
			if seen[sub] {
				res = sub
				return true
			}
			seen[sub] = true
		}
		return false
	}

	l, r := 0, len(s)
	for l < r {
		m := r - (r+l)/2
		if isExist(m) {
			l = m
		} else {
			r = m - 1
		}
	}

	return res
}
