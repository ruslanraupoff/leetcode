package main

import "strconv"

func decodeString(s string) string {
	r, _ := dfs(0, s)
	return r
}

func dfs(i int, s string) (string, int) {
	c, r, t := "", "", ""
	for i < len(s) {
		if s[i] == '[' {
			t, i = dfs(i+1, s)
			a, _ := strconv.Atoi(c)
			r += times(a, t)
			c = ""
		} else if s[i] == ']' {
			return r, i
		} else if '0' <= s[i] && s[i] <= '9' {
			c += string(s[i])
		} else {
			r += string(s[i])
		}
		i++
	}
	return r, i
}

func times(n int, t string) string {
	r := ""
	for i := 0; i < n; i++ {
		r += t
	}
	return r
}
